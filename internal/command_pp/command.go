package commandpp

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"homework/internal/command_pp/request"
	"homework/internal/command_pp/response"
	"homework/internal/pick-up_point/delivery"
)

type Commands []Command

func InitCommands(
	ppDelivery *delivery.PPDelivery,
	chanForRead chan<- request.Request,
	chanForWrite chan<- request.Request,
) Commands {
	return Commands{
		New("add", chanForWrite, ppDelivery.AddPickUpPoint),
		New("get_by_name", chanForRead, ppDelivery.GetPickUpPointByName),
		New("get_all", chanForRead, ppDelivery.GetPickUpPoints),
		New("update", chanForWrite, ppDelivery.UpdatePickUpPoint),
	}
}

func (cs Commands) Call(commandName string, params []string) (uuid.UUID, error) {
	for _, cmd := range cs {
		if cmd.Name == commandName {
			req := request.NewRequest(cmd.FuncToCall, params)
			cmd.RequestChan <- req
			return req.ID, nil
		}
	}
	return uuid.UUID{}, fmt.Errorf("unknown command")
}

func (cs Commands) ProcessInput(ctx context.Context) {
	inputChan := make(chan string)
	go input(inputChan)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("End receiving commands")
			return
		case paramsStr := <-inputChan:
			params := strings.Split(paramsStr, " ")
			if len(params) == 0 {
				fmt.Println("User info: no command entered")
				continue
			}
			reqID, err := cs.Call(params[0], params[1:])
			if err != nil {
				fmt.Printf("User info: ended with error: %s\n", err)
			} else {
				fmt.Printf("User info: your request got ID %s\n", reqID)
			}
		}
	}
}

func input(inputChan chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("User info: Waiting for command:")
		scanner.Scan()
		paramsStr := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("User info:Error in reading input:", err)
		}
		inputChan <- paramsStr
	}
}

func ProcessResponses(responseChan <-chan response.Response) {
	for resp := range responseChan {
		if resp.Err != nil {
			fmt.Printf("User info: error for request %s: %v\n", resp.ID, resp.Err)
		} else {
			fmt.Printf("User info: success for request %s: %s\n", resp.ID, resp.Body)
		}
	}
}

func ProcessLogs(logChan <-chan string) {
	for newLog := range logChan {
		fmt.Printf("Log info: %s\n", newLog)
	}
}
