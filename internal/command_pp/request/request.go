package request

import (
	"context"

	"github.com/google/uuid"
	"homework/internal/command_pp/response"
)

type handlerFn func(context.Context, []string) response.Response

type Request struct {
	ID         uuid.UUID
	FuncToCall handlerFn
	Params     []string
}

func NewRequest(funcToCall handlerFn, params []string) Request {
	return Request{
		ID:         uuid.New(),
		FuncToCall: funcToCall,
		Params:     params,
	}
}
