package dto

import (
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/model/address"
)

func ConvertPPAddToPickUpPoint(p PickUpPointAdd) model.PickUpPoint {
	return model.PickUpPoint{
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

func ConvertPPUpdateToPickUpPoint(p PickUpPointUpdate) model.PickUpPoint {
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
