package service

import "time"

func (op *OrderServicePP) DeleteOrderService(orderID int) error {
	orders, err := op.storage.GetOrders()
	if err != nil {
		return err
	}
	for _, order := range orders {
		if order.ID != orderID {
			continue
		}
		if order.IsReturned {
			return op.storage.DeleteOrderStorage(orderID)
		}
		if order.IsIssued {
			return ErrOrderAlreadyIssued
		}
		if order.StorageExpirationDate.After(time.Now()) {
			return ErrOrderShelfLifeNotExpired
		}
		return op.storage.DeleteOrderStorage(orderID)
	}
	return ErrOrderNotFound
}
