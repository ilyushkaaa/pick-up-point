package dto

import (
	"homework/internal/pick-up_point/model"
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

func NewPickUpPointDB(p model.PickUpPoint) PickUpPointDB {
	return PickUpPointDB{
		ID:          p.ID,
		Name:        p.Name,
		PhoneNumber: p.PhoneNumber,
		Region:      p.Address.Region,
		City:        p.Address.City,
		Street:      p.Address.Street,
		HouseNum:    p.Address.HouseNum,
	}
}
