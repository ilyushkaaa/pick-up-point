package model

type Command struct {
	Name        string
	Params      []string
	Description string
	FuncToCall  func([]string)
}

func NewCommand(name, description string, params []string, funcToCall func([]string)) Command {
	return Command{
		Name:        name,
		Description: description,
		Params:      params,
		FuncToCall:  funcToCall,
	}
}
