package dto

import (
	"github.com/asaskevich/govalidator"
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/model/address"
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

func (p *PickUpPointUpdate) ConvertToPickUpPoint() model.PickUpPoint {
	return model.PickUpPoint{
		ID:   p.ID,
		Name: p.Name,
		Address: address.PPAddress{
			Region:   p.Address.Region,
			City:     p.Address.City,
			Street:   p.Address.Street,
			HouseNum: p.Address.HouseNum,
		},
		PhoneNumber: p.PhoneNumber,
	}
}
