package delivery

import (
	"encoding/json"
	"fmt"

	"homework/internal/pick-up_point/model"
)

func (ps *PPDelivery) Validate(pickUpPoint model.PickUpPoint) error {
	if validationErrors := pickUpPoint.Validate(); len(validationErrors) != 0 {
		var errorsJSON []byte
		errorsJSON, err := json.Marshal(validationErrors)
		if err != nil {

			return fmt.Errorf("you have problems with validation, but there was an error in coding them in json: %w", err)

		}
		return fmt.Errorf("validation errors: %s", string(errorsJSON))

	}
	return nil
}
