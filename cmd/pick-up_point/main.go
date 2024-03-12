package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
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

	chanForWrite := make(chan request.Request)
	chanForRead := make(chan request.Request)
	responseChan := make(chan response.Response)
	logChan := make(chan string)
	inputChan := make(chan string)

	PPStorage, err := storage.New(logChan)
	if err != nil {
		fmt.Printf("error in creating storage: %s\n", err)
		return
	}

	wg := &sync.WaitGroup{}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer func() {
		cancel()
		wg.Wait()
		err = PPStorage.Close()
		if err != nil {
			fmt.Printf("error in closing storage: %s\n", err)
		} else {
			fmt.Println("storage was closed successfully")
		}
	}()

	PPService := service.New(PPStorage)
	PPDelivery := delivery.New(PPService)

	commands := commandpp.InitCommands(PPDelivery, chanForRead, chanForWrite)

	go worker.Work(chanForWrite, responseChan, logChan)
	go worker.Work(chanForRead, responseChan, logChan)

	go commandpp.ProcessResponses(responseChan)
	go func() {
		wg.Add(1)
		defer wg.Done()
		commands.ProcessInput(ctx, inputChan)
	}()
	go commandpp.ProcessLogs(logChan)

	<-sgChan
	fmt.Println("Got end signal")
}
