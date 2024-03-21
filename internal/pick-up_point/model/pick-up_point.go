package model

import (
	"homework/internal/pick-up_point/model/address"
)

type PickUpPoint struct {
	ID          uint64
	Name        string
	Address     address.PPAddress
	PhoneNumber string
}
