package main

import (
	"fmt"
	"os"

	"homework/internal/command"
	"homework/internal/order/delivery"
	"homework/internal/order/service"
	"homework/internal/order/storage"
)

const tipText = `To see full list of commands start program with command line argument "help"`

func main() {
	orderStorage, err := storage.New()
	if err != nil {
		fmt.Printf("error in creating storage: %s\n", err)
		return
	}
	defer func() {
		err = orderStorage.Close()
		if err != nil {
			fmt.Printf("error in closing storage: %s", err)
		}
	}()
	orderService := service.New(orderStorage)
	orderDelivery := delivery.New(orderService)
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Command is not set: %s\n", tipText)
		return
	}
	commands := command.InitCommands(orderDelivery)
	err = commands.Call(args[0], args[1:])
	if err != nil {
		fmt.Printf("ended with error: %s", err)
	}
}
