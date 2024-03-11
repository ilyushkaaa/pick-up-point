package delivery

import (
	"encoding/json"
	"fmt"

	"homework/internal/command_pp/response"
	"homework/internal/pick-up_point/model"
)

func (ps *PPDelivery) AddPickUpPointDelivery(params []string) response.Response {
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

	if validationErrors := newPickUpPoint.Validate(); len(validationErrors) != 0 {
		var errorsJSON []byte
		errorsJSON, err = json.Marshal(validationErrors)
		if err != nil {
			return response.Response{
				Err: fmt.Errorf("you have problems with validation, but there was an error in coding them in json: %w", err),
			}
		}
		return response.Response{
			Err: fmt.Errorf("validation errors: %s", string(errorsJSON)),
		}
	}
	err = ps.service.AddPickUpPointService(newPickUpPoint)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("error in adding new pick-up point: %w", err),
		}
	}
	return response.Response{
		Body: "pick-up point was successfully added",
	}
}
