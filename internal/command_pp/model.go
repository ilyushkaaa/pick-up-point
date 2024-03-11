package commandpp

import (
	"homework/internal/command_pp/request"
	"homework/internal/command_pp/response"
)

type Command struct {
	Name        string
	FuncToCall  func([]string) response.Response
	RequestChan chan<- request.Request
}

func New(name string, requestChan chan<- request.Request, funcToCall func([]string) response.Response) Command {
	return Command{
		Name:        name,
		FuncToCall:  funcToCall,
		RequestChan: requestChan,
	}
}
