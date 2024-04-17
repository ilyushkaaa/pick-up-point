package dto

import "homework/internal/order/model"

func ConvertToOrder(o OrderFromCourierInputData) model.Order {
	return model.Order{
		ID:                    o.ID,
		ClientID:              o.ClientID,
		Weight:                o.Weight,
		Price:                 o.Price,
		StorageExpirationDate: o.StorageExpirationDate,
		PackageType:           o.PackageType,
		PickUpPointID:         o.PickUpPointID,
	}
}
