package delivery

import (
	"context"
	"encoding/json"
	"fmt"

	"homework/internal/command_pp/response"
	"homework/internal/pick-up_point/delivery/dto"
)

func (ps *PPDelivery) AddPickUpPoint(ctx context.Context, params []string) response.Response {
	if len(params) != 1 {
		return response.Response{
			Err: fmt.Errorf("add pick-up point method must have 1 param"),
		}
	}
	var pp dto.PickUpPointAdd
	err := json.Unmarshal([]byte(params[0]), &pp)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("new pick-up point must be valid json: %w", err),
		}
	}
	err = pp.Validate()
	if err != nil {
		return response.Response{
			Err: err,
		}
	}
	newPickUpPoint := dto.ConvertPPAddToPickUpPoint(pp)
	addedPP, err := ps.service.AddPickUpPoint(ctx, newPickUpPoint)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("error in adding new pick-up point: %w", err),
		}
	}
	return response.Response{
		Body: fmt.Sprintf("pick-up point was successfully added: %v", addedPP),
	}
}
