package delivery

import (
	"fmt"

	"homework/internal/command_pp/response"
)

func (ps *PPDelivery) GetPickUpPointByName(input []string) response.Response {
	if len(input) != 1 {
		return response.Response{
			Err: fmt.Errorf("bad input params number: it must be 1"),
		}
	}
	nameLen := len(input[0])
	if nameLen < 5 || nameLen > 50 {
		return response.Response{
			Err: fmt.Errorf("pick-up point name must contain from 5 to 50 symbols"),
		}
	}
	pickUpPoint, err := ps.service.GetPickUpPointByName(input[0])
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("error in getting pick-up pointby name: %w", err),
		}
	}
	return response.Response{
		Body: fmt.Sprintf("%+v", pickUpPoint),
	}
}
