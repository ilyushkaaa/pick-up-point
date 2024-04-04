package delivery

import (
	"context"
	"encoding/json"
	"fmt"

	"homework/internal/command_pp/response"
	"homework/internal/pick-up_point/delivery/dto"
)

func (ps *PPDelivery) UpdatePickUpPoint(ctx context.Context, params []string) response.Response {
	if len(params) != 1 {
		return response.Response{
			Err: fmt.Errorf("update pick-up point method must have 1 param"),
		}
	}
	var pp dto.PickUpPointUpdate
	err := json.Unmarshal([]byte(params[0]), &pp)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("potential updated pick-up point must be valid json: %w", err),
		}
	}
	err = pp.Validate()
	if err != nil {
		return response.Response{
			Err: err,
		}
	}
	pickUpPoint := dto.ConvertPPUpdateToPickUpPoint(pp)
	err = ps.service.UpdatePickUpPoint(ctx, pickUpPoint)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("error in updating pick-up point: %w", err),
		}
	}
	return response.Response{
		Body: "pick-up point was successfully updated",
	}
}
