package dto

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type OrderFromCourierInputData struct {
	ID                    uint64    `json:"id" valid:"required"`
	ClientID              uint64    `json:"client_id" valid:"required"`
	Weight                float64   `json:"weight" valid:"required"`
	Price                 float64   `json:"price" valid:"required"`
	StorageExpirationDate time.Time `json:"storage_expiration_date" valid:"required"`
	PackageType           string    `json:"package_type"`
	PickUpPointID         uint64    `json:"pick_up_point_id" valid:"required"`
}

func (o *OrderFromCourierInputData) Validate() error {
	_, err := govalidator.ValidateStruct(o)
	return err
}
