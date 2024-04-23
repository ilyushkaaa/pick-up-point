package dto

import (
	"github.com/asaskevich/govalidator"
)

type PickUpPointAdd struct {
	Name        string     `json:"name" valid:"required,length(5|50)"`
	Address     AddressDTO `json:"address" valid:"required"`
	PhoneNumber string     `json:"phone_number" valid:"required,matches(^[0-9]+$)"`
}

func (p *PickUpPointAdd) Validate() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
