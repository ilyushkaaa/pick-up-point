package dto

import (
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/model/address"
)

func ConvertToPickUpPoint(p PickUpPointDB) model.PickUpPoint {
	return model.PickUpPoint{
		ID:          p.ID,
		Name:        p.Name,
		PhoneNumber: p.PhoneNumber,
		Address: address.PPAddress{
			Region:   p.Region,
			City:     p.City,
			Street:   p.Street,
			HouseNum: p.HouseNum,
		},
	}
}

func ConvertSliceToPickUpPoints(p []PickUpPointDB) []model.PickUpPoint {
	pickUpPoints := make([]model.PickUpPoint, 0, len(p))
	for _, pp := range p {
		pickUpPoints = append(pickUpPoints, ConvertToPickUpPoint(pp))
	}
	return pickUpPoints
}
