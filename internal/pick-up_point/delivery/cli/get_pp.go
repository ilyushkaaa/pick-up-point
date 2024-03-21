package delivery

import (
	"context"
	"fmt"

	"homework/internal/command_pp/response"
)

func (ps *PPDelivery) GetPickUpPoints(ctx context.Context, input []string) response.Response {
	if len(input) != 0 {
		return response.Response{
			Err: fmt.Errorf("this request must not contain any params"),
		}
	}
	pickUpPoints, err := ps.service.GetPickUpPoints(ctx)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("error in getting pick-up points: %w", err),
		}
	}
	return response.Response{
		Body: fmt.Sprintf("%+v", pickUpPoints),
	}
}
