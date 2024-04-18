package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	cacheInMemory "homework/internal/cache/in_memory"
	cacheRedis "homework/internal/cache/redis"
	"homework/internal/middleware"
	deliveryOrder "homework/internal/order/delivery/http"
	serviceOrder "homework/internal/order/service"
	"homework/internal/order/service/packages"
	storageOrder "homework/internal/order/storage/database"
	deliveryPP "homework/internal/pick-up_point/delivery/http"
	servicePP "homework/internal/pick-up_point/service"
	storagePP "homework/internal/pick-up_point/storage/database"
	"homework/internal/routes"
	database "homework/pkg/database/postgres"
	"homework/pkg/database/postgres/transaction_manager"
	"homework/pkg/kafka"
	"homework/pkg/kafka/consumer"
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
	err = godotenv.Load(".env")
	if err != nil {
		logger.Fatalf("error in getting env: %s", err)
	}
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
	redisCache := cacheRedis.New(logger)
	defer func() {
		err = redisCache.Close()
		if err != nil {
			logger.Errorf("error in closing redis cache: %v", err)
		}
	}()

	imMemoryCache := cacheInMemory.New()

	stPP := storagePP.New(db)
	stOrder := storageOrder.New(db)

	svPP := servicePP.New(stPP, stOrder, tm, imMemoryCache)
	dPP := deliveryPP.New(redisCache, svPP, logger)

	packageTypes := packages.Init()
	svOrder := serviceOrder.New(stOrder, stPP, packageTypes, tm, imMemoryCache)
	dOrder := deliveryOrder.New(svOrder, logger)

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

	mw := middleware.New(logger, cfg.Producer)
	router := routes.GetRouter(dPP, dOrder, mw)

	waitChan := make(chan struct{})

	consumer.GoRunConsumer(ctx, cfg, logger, waitChan)
	err = consumer.WaitForConsumerReady(waitChan)
	if err != nil {
		logger.Fatal(err)
	}

	port := os.Getenv("APP_PORT")
	addr := ":" + port
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)

	logger.Fatal(http.ListenAndServeTLS(addr, "./server.crt", "./server.key", router))
}
