package service

import (
	"time"

	"homework/internal/order/model"
)

func (op *OrderServicePP) AddOrderService(orderID, clientID int, expireDate time.Time) error {
	orders, err := op.storage.GetOrders()
	if err != nil {
		return err
	}
	for _, order := range orders {
		if order.ID == orderID {
			return ErrOrderAlreadyExists
		}
	}
	if time.Now().After(expireDate) {
		return ErrShelfTimeExpired
	}
	newOrder := model.Order{
		ID:                    orderID,
		ClientID:              clientID,
		StorageExpirationDate: expireDate,
		OrderIssueDate:        time.Time{},
		IsIssued:              false,
		IsReturned:            false,
	}
	return op.storage.AddOrderStorage(newOrder)
}
