package worker

import (
	"context"
	"fmt"

	"homework/internal/command_pp/request"
	"homework/internal/command_pp/response"
)

func Work(ctx context.Context, requestChan <-chan request.Request, responseChan chan<- response.Response, logChan chan<- string) {
	for req := range requestChan {
		logChan <- fmt.Sprintf("New request with ID %s was received: %v", req.ID, req)
		resp := req.FuncToCall(ctx, req.Params)
		resp.ID = req.ID
		logChan <- fmt.Sprintf("Response for request %s was received: %v", req.ID, resp)
		responseChan <- resp
		logChan <- fmt.Sprintf("Response %s was sent for user", resp.ID)
	}
}
