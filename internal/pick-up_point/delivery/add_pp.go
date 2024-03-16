package delivery

import (
	"encoding/json"
	"fmt"

	"homework/internal/command_pp/response"
	"homework/internal/pick-up_point/model"
)

func (ps *PPDelivery) AddPickUpPoint(params []string) response.Response {
	if len(params) != 1 {
		return response.Response{
			Err: fmt.Errorf("add pick-up point method must have 1 param"),
		}
	}
	var newPickUpPoint model.PickUpPoint
	err := json.Unmarshal([]byte(params[0]), &newPickUpPoint)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("new pick-up point must be valid json: %w", err),
		}
	}

	if err = ps.Validate(newPickUpPoint); err != nil {
		return response.Response{
			Err: err,
		}
	}
	err = ps.service.AddPickUpPoint(newPickUpPoint)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("error in adding new pick-up point: %w", err),
		}
	}
	return response.Response{
		Body: "pick-up point was successfully added",
	}
}
