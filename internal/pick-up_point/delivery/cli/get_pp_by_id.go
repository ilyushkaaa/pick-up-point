package delivery

import (
	"context"
	"fmt"
	"strconv"

	"homework/internal/command_pp/response"
)

func (ps *PPDelivery) GetPickUpPointByID(ctx context.Context, input []string) response.Response {
	if len(input) != 1 {
		return response.Response{
			Err: fmt.Errorf("bad input params number: it must be 1"),
		}
	}
	ppID, err := strconv.ParseUint(input[0], 10, 64)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("pick-up point ID %q must be positive integer", input[0]),
		}
	}
	pickUpPoint, err := ps.service.GetPickUpPointByID(ctx, ppID)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("error in getting pick-up point by id: %w", err),
		}
	}
	return response.Response{
		Body: fmt.Sprintf("%+v", pickUpPoint),
	}
}
