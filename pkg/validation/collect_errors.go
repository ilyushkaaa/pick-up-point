package validation

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

func CollectErrors(err error) error {
	validationErrors := make([]string, 0)
	if err == nil {
		return nil
	}
	var allErrs govalidator.Errors
	if errors.As(err, &allErrs) {
		for _, fld := range allErrs {
			validationErrors = append(validationErrors, fld.Error())
		}
	}
	errorsJSON, err := json.Marshal(validationErrors)
	if err != nil {
		return fmt.Errorf("you have problems with validation, but there was an error in coding them in json: %w", err)
	}
	return fmt.Errorf("validation errors: %s", string(errorsJSON))
}
