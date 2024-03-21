package request

import (
	"context"

	"github.com/google/uuid"
	"homework/internal/command_pp/response"
)

type Request struct {
	ID         uuid.UUID
	FuncToCall func(context.Context, []string) response.Response
	Params     []string
}

func NewRequest(funcToCall func(context.Context, []string) response.Response, params []string) Request {
	return Request{
		ID:         uuid.New(),
		FuncToCall: funcToCall,
		Params:     params,
	}
}
