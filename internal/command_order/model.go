package commandorder

import "context"

type CallFunc func(context.Context, []string) error

type Command struct {
	Name        string
	Params      []string
	Description string
	FuncToCall  CallFunc
}

func New(name, description string, params []string, funcToCall CallFunc) Command {
	return Command{
		Name:        name,
		Description: description,
		Params:      params,
		FuncToCall:  funcToCall,
	}
}
