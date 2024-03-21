package dto

import (
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/model/address"
)

type PickUpPointDB struct {
	ID          uint64 `db:"id"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
	Region      string `db:"region"`
	City        string `db:"city"`
	Street      string `db:"street"`
	HouseNum    string `db:"house_num"`
}

func (pp *PickUpPointDB) ConvertToPickUpPoint() model.PickUpPoint {
	return model.PickUpPoint{
		ID:          pp.ID,
		Name:        pp.Name,
		PhoneNumber: pp.PhoneNumber,
		Address: address.PPAddress{
			Region:   pp.Region,
			City:     pp.City,
			Street:   pp.Street,
			HouseNum: pp.HouseNum,
		},
	}
}
