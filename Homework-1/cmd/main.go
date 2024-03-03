package main

import (
	"fmt"
	"os"

	"homework/Homework-1/internal/command"
	"homework/Homework-1/internal/order/delivery"
	"homework/Homework-1/internal/order/service"
	"homework/Homework-1/internal/order/storage"
)

const tipText = `To see full list of commands start program with command line argument "help"`

func main() {
	fileName := "storage.json"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error in open file: %s\n", err)
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("error in closing file: %s\n", err)
		}
	}(file)
	fileOrderStorage := storage.NewFileOrderStorage(file)
	fileOrderService := service.NewOrderServicePP(fileOrderStorage)
	fileOrderDelivery := delivery.NewOrderDelivery(fileOrderService)
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Command is not set: %s\n", tipText)
		return
	}
	commands := command.InitCommands(fileOrderDelivery)
	switch args[0] {
	case "help":
		command.PrintCommandsDescription(commands)
	default:
		for _, cmd := range commands {
			if cmd.Name == args[0] {
				cmd.FuncToCall(args[1:])
				return
			}
		}
		fmt.Printf("Unknown command: %s\n", tipText)
	}

}
