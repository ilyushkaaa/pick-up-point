package model

import (
	"errors"

	"homework/internal/pick-up_point/model/address"

	"github.com/asaskevich/govalidator"
)

type PickUpPoint struct {
	Name        string            `valid:"required,length(5|50)"`
	Address     address.PPAddress `valid:"required"`
	PhoneNumber string            `valid:"required,matches(^0-9]+$)"`
}

func (u *PickUpPoint) Validate() []string {
	_, err := govalidator.ValidateStruct(u)
	return collectErrors(err)
}

func collectErrors(err error) []string {
	validationErrors := make([]string, 0)
	if err == nil {
		return validationErrors
	}
	var allErrs govalidator.Errors
	if errors.As(err, &allErrs) {
		for _, fld := range allErrs {
			validationErrors = append(validationErrors, fld.Error())
		}
	}
	return validationErrors
}
