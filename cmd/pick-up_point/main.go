package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"homework/internal/command_pp"
	"homework/internal/command_pp/request"
	"homework/internal/command_pp/response"
	"homework/internal/pick-up_point/delivery"
	"homework/internal/pick-up_point/service"
	"homework/internal/pick-up_point/storage"
	"homework/internal/pick-up_point/worker"
)

func main() {
	sgChan := make(chan os.Signal, 1)
	signal.Notify(sgChan, syscall.SIGINT, syscall.SIGTERM)

	PPStorage, err := storage.New()
	if err != nil {
		fmt.Printf("error in creating storage: %s\n", err)
		return
	}
	defer func() {
		fmt.Println("was closed")
		err = PPStorage.Close()
		if err != nil {
			fmt.Printf("error in closing storage: %s\n", err)
		}
	}()

	PPService := service.New(PPStorage)
	PPDelivery := delivery.New(PPService)

	chanForWrite := make(chan request.Request)
	chanForRead := make(chan request.Request)
	responseChan := make(chan response.Response)
	logChan := make(chan string)

	commands := commandpp.InitCommands(PPDelivery, chanForRead, chanForWrite)

	go worker.Work(chanForWrite, responseChan, logChan)
	go worker.Work(chanForRead, responseChan, logChan)

	go commandpp.ProcessResponses(responseChan)
	go commands.ProcessInput()
	go commandpp.ProcessLogs(logChan)

	<-sgChan
	fmt.Println("Got end signal")
}
