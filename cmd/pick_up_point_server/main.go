package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"homework/internal/cache"
	cacheInMemory "homework/internal/cache/in_memory"
	cacheRedis "homework/internal/cache/redis"
	"homework/internal/interceptor"
	deliveryOrder "homework/internal/order/delivery/grpc"
	serviceOrder "homework/internal/order/service"
	"homework/internal/order/service/packages"
	storageOrder "homework/internal/order/storage/database"
	pbOrder "homework/internal/pb/order"
	pbPP "homework/internal/pb/pick-up_point"
	deliveryPP "homework/internal/pick-up_point/delivery/grpc"
	servicePP "homework/internal/pick-up_point/service"
	storagePP "homework/internal/pick-up_point/storage/database"
	database "homework/pkg/infrastructure/database/postgres"
	"homework/pkg/infrastructure/database/postgres/transaction_manager"
	"homework/pkg/infrastructure/jaeger"
	"homework/pkg/infrastructure/kafka"
	"homework/pkg/infrastructure/kafka/consumer"
	"homework/pkg/infrastructure/prometheus"
)

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("error in logger initialization: %v", err)
	}
	logger := zapLogger.Sugar()
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Printf("error in logger sync: %v", err)
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tm, err := transaction_manager.New(ctx)
	if err != nil {
		logger.Fatalf("error in creating transaction manager: %v", err)
	}

	db := database.New(tm)
	if err != nil {
		logger.Fatalf("error in database init: %v", err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			logger.Errorf("error in closing db")
		}
	}()

	cacheConfig, err := cache.GetConfig()
	if err != nil {
		logger.Fatalf("error in getting cache config: %v", err)
	}

	redisCache := cacheRedis.New(logger, cacheConfig.RedisAddr, cacheConfig.RedisPassword, cacheConfig.RedisTTl)
	defer func() {
		err = redisCache.Close()
		if err != nil {
			logger.Errorf("error in closing redis cache: %v", err)
		}
	}()

	cfg, err := kafka.NewConfig()
	if err != nil {
		logger.Fatalf("error in kafka config init: %v", err)
	}
	defer func() {
		err = cfg.Close()
		if err != nil {
			logger.Errorf("error in closing sync kafka producer: %v", err)
		}
	}()

	shutdown, err := jaeger.InitProvider()
	if err != nil {
		logger.Fatalf("error in tracer init: %v", err)
	}
	defer func() {
		if err = shutdown(ctx); err != nil {
			logger.Fatalf("failed to shutdown TracerProvider: %v", err)
		}
	}()

	infra := infrastructure{
		tm:          tm,
		db:          db,
		cacheConfig: cacheConfig,
		redisCache:  redisCache,
		cfg:         cfg,
	}
	goRunGRPCServer(ctx, infra, logger)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = pbOrder.RegisterOrdersHandlerFromEndpoint(ctx, mux, "localhost:9011", opts)
	if err != nil {
		logger.Fatalf("failed to register grpc gateway order handler: %v", err)
	}
	err = pbPP.RegisterPickUpPointsHandlerFromEndpoint(ctx, mux, "localhost:9011", opts)
	if err != nil {
		logger.Fatalf("failed to register grpc gateway pick-up points handler: %v", err)
	}

	port := os.Getenv("APP_PORT")
	addr := ":" + port
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)

	logger.Fatal(http.ListenAndServeTLS(addr, "./server.crt", "./server.key", mux))
}

func goRunGRPCServer(ctx context.Context, infra infrastructure, logger *zap.SugaredLogger) {
	lis, err := net.Listen("tcp", ":9011")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tracer := otel.Tracer("test-tracer")

	i := interceptor.New(logger, infra.cfg.Producer)
	grpcMetrics := grpc_prometheus.NewServerMetrics()

	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(grpcMetrics.UnaryServerInterceptor(), i.AccessLog, i.Metric, i.Auth),
	)

	bm, sm := prometheus.Init(s, logger, grpcMetrics)

	i.SetMetrics(sm)

	orderByIDCache := cacheInMemory.New(logger, infra.cacheConfig.OrderByIDTTl, infra.cacheConfig.Capacity)
	ppByIDCache := cacheInMemory.New(logger, infra.cacheConfig.PPByIDTTl, infra.cacheConfig.Capacity)
	ordersByClientCache := cacheInMemory.New(logger, infra.cacheConfig.OrdersByClientTTl, infra.cacheConfig.Capacity)

	stPP := storagePP.New(infra.db, tracer)
	stOrder := storageOrder.New(infra.db, tracer)

	svPP := servicePP.New(stPP, stOrder, infra.tm, ppByIDCache, tracer)

	packageTypes := packages.Init()
	svOrder := serviceOrder.New(stOrder, stPP, packageTypes, infra.tm, orderByIDCache, ordersByClientCache, ppByIDCache, bm, tracer)

	waitChan := make(chan struct{})

	consumer.GoRunConsumer(ctx, infra.cfg, logger, waitChan)
	err = consumer.WaitForConsumerReady(waitChan)
	if err != nil {
		logger.Fatal(err)
	}

	pbOrder.RegisterOrdersServer(s, deliveryOrder.New(svOrder, logger, tracer))
	pbPP.RegisterPickUpPointsServer(s, deliveryPP.New(infra.redisCache, svPP, logger, tracer))

	logger.Infow("starting grpc server",
		"type", "START",
		"addr", ":9011",
	)
	go func() {
		logger.Fatal(s.Serve(lis))
	}()
}

type infrastructure struct {
	tm          transaction_manager.TransactionManager
	db          database.Database
	cacheConfig *cache.Config
	redisCache  cache.Cache
	cfg         *kafka.ConfigKafka
}
