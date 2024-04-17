package dto

import (
	"homework/internal/order/model"
)

func ConvertToOrder(o OrderDB) model.Order {
	order := model.Order{
		ID:                    o.ID,
		ClientID:              o.ClientID,
		Weight:                o.Weight,
		Price:                 o.Price,
		StorageExpirationDate: o.StorageExpirationDate,
		IsReturned:            o.IsReturned,
		PickUpPointID:         o.PickUpPointID,
	}
	if o.OrderIssueDate.Valid {
		order.OrderIssueDate = o.OrderIssueDate.Time
		order.IsIssued = true
	}
	if o.PackageType.Valid {
		order.PackageType = o.PackageType.String
	}
	return order
}

func ConvertSliceToOrders(o []OrderDB) []model.Order {
	orders := make([]model.Order, 0, len(o))
	for _, order := range o {
		orders = append(orders, ConvertToOrder(order))
	}
	return orders
}
