package commandorder

type Command struct {
	Name        string
	Params      []string
	Description string
	FuncToCall  func([]string) error
}

func New(name, description string, params []string, funcToCall func([]string) error) Command {
	return Command{
		Name:        name,
		Description: description,
		Params:      params,
		FuncToCall:  funcToCall,
	}
}
