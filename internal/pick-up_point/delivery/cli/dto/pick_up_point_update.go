package dto

import (
	"github.com/asaskevich/govalidator"
)

type PickUpPointUpdate struct {
	ID          uint64     `json:"ID" valid:"required"`
	Name        string     `json:"name" valid:"required,length(5|50)"`
	Address     AddressDTO `json:"address" valid:"required"`
	PhoneNumber string     `json:"phone_number" valid:"required,matches(^[0-9]+$)"`
}

func (p *PickUpPointUpdate) Validate() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
