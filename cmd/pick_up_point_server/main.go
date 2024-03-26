package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
	"homework/internal/middleware"
	delivery "homework/internal/pick-up_point/delivery/http"
	"homework/internal/pick-up_point/delivery/http/routes"
	"homework/internal/pick-up_point/service"
	storage "homework/internal/pick-up_point/storage/database"
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
	st, err := storage.New(ctx)
	if err != nil {
		logger.Fatalf("error in storage init: %v", err)
	}
	sv := service.New(st)
	d := delivery.New(sv, logger)
	mw := middleware.New(logger)
	router := routes.GetRouter(d, mw)

	port := os.Getenv("APP_PORT")
	addr := ":" + port
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)
	logger.Fatal(http.ListenAndServeTLS(addr, "./server.crt", "./server.key", router))
}
