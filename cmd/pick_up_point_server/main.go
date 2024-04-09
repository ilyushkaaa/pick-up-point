package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
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
)

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("error in events initialization: %v", err)
	}
	logger := zapLogger.Sugar()
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Printf("error in events sync: %v", err)
		}
	}()
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

	mw := middleware.New(logger)
	router := routes.GetRouter(dPP, dOrder, mw)

	port := os.Getenv("APP_PORT")
	addr := ":" + port
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)
	logger.Fatal(http.ListenAndServeTLS(addr, "./server.crt", "./server.key", router))
}
