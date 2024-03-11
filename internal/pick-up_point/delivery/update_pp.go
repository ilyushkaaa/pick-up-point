package delivery

import (
	"encoding/json"
	"fmt"

	"homework/internal/command_pp/response"
	"homework/internal/pick-up_point/model"
)

func (ps *PPDelivery) UpdatePickUpPoint(params []string) response.Response {
	if len(params) != 1 {
		return response.Response{
			Err: fmt.Errorf("update pick-up point method must have 1 param"),
		}
	}
	var pickUpPoint model.PickUpPoint
	err := json.Unmarshal([]byte(params[0]), &pickUpPoint)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("potential updated pick-up point must be valid json: %w", err),
		}
	}

	if err = ps.Validate(pickUpPoint); err != nil {
		return response.Response{
			Err: err,
		}
	}

	err = ps.service.UpdatePickUpPoint(pickUpPoint)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("error in updating pick-up point: %w", err),
		}
	}
	return response.Response{
		Body: "pick-up point was successfully updated",
	}
}
