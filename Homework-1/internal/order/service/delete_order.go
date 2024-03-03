package service

import (
	"time"

	"homework/Homework-1/internal/order/myerrors"
)

func (op *OrderServicePP) DeleteOrderService(orderID int) error {
	orders, err := op.storage.GetOrders()
	if err != nil {
		return err
	}
	for _, ord := range orders {
		if ord.ID != orderID {
			continue
		}
		if ord.IsReturned {
			return op.storage.DeleteOrderStorage(orderID)
		}
		if ord.IsIssued {
			return myerrors.ErrorOrderAlreadyIssued
		}
		if ord.StorageExpirationDate.After(time.Now()) {
			return myerrors.ErrorOrderShelfLifeNotExpired
		}
		return op.storage.DeleteOrderStorage(orderID)
	}
	return myerrors.ErrorOrderNotFound
}
