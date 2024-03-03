package service

import (
	"time"

	"homework/Homework-1/internal/order/model"
	"homework/Homework-1/internal/order/myerrors"
)

func (op *OrderServicePP) AddOrderService(orderID, clientID int, expireDate time.Time) error {
	orders, err := op.storage.GetOrders()
	if err != nil {
		return err
	}
	for _, ord := range orders {
		if ord.ID == orderID {
			return myerrors.ErrorOrderAlreadyExists
		}
	}
	if time.Now().After(expireDate) {
		return myerrors.ErrorShelfTimeExpired
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
