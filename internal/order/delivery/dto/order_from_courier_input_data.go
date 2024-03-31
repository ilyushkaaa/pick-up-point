package dto

import (
	"time"

	"github.com/asaskevich/govalidator"
	"homework/internal/order/model"
)

type OrderFromCourierInputData struct {
	ID                    uint64    `json:"id" valid:"required"`
	ClientID              uint64    `json:"client_id" valid:"required"`
	Weight                float64   `json:"weight" valid:"required"`
	Price                 float64   `json:"price" valid:"required"`
	StorageExpirationDate time.Time `json:"storage_expiration_date" valid:"required"`
	PackageType           string    `json:"package_type"`
}

func (o *OrderFromCourierInputData) Validate() error {
	_, err := govalidator.ValidateStruct(o)
	return err
}

func (o *OrderFromCourierInputData) ConvertToOrder() model.Order {
	return model.Order{
		ID:                    o.ID,
		ClientID:              o.ClientID,
		Weight:                o.Weight,
		Price:                 o.Price,
		StorageExpirationDate: o.StorageExpirationDate,
		PackageType:           o.PackageType,
	}
}
