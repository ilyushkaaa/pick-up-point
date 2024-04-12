package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	eventsProducer "homework/internal/events/service/producer"
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
	"homework/pkg/kafka/consumer"
	"homework/pkg/kafka/producer"
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
	db, err := database.New(ctx)
	if err != nil {
		logger.Fatalf("error in database init: %v", err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			logger.Errorf("error in closing db")
		}
	}()
	stPP := storagePP.New(db)
	svPP := servicePP.New(stPP)
	dPP := deliveryPP.New(svPP, logger)

	stOrder := storageOrder.New(db)

	packageTypes := packages.Init()
	svOrder := serviceOrder.New(stOrder, packageTypes)
	dOrder := deliveryOrder.New(svOrder, logger)

	brokers := []string{os.Getenv("KAFKA1_ADDR"), os.Getenv("KAFKA2_ADDR"), os.Getenv("KAFKA3_ADDR")}
	syncProducer, err := producer.New(brokers)
	if err != nil {
		logger.Fatalf("error in kafka producer create: %s", err)
	}

	defer func() {
		err = syncProducer.Close()
		if err != nil {
			logger.Errorf("error in closing sync kafka producer: %s", err)
		}
	}()

	topic := os.Getenv("KAFKA_EVENTS_TOPIC")
	groupID := os.Getenv("EVENTS_CONSUMER_GROUP_ID")
	ep := eventsProducer.New(syncProducer, topic)
	mw := middleware.New(logger, ep)
	router := routes.GetRouter(dPP, dOrder, mw)

	go func() {
		err = consumer.Run(brokers, logger, ctx, topic, groupID)
		if err != nil {
			logger.Errorf("error in consumer running")
		}
	}()

	port := os.Getenv("APP_PORT")
	addr := ":" + port
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)

	logger.Fatal(http.ListenAndServeTLS(addr, "./server.crt", "./server.key", router))
}
