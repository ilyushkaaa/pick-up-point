package service

import "time"

func (op *OrderServicePP) ReturnOrderService(clientID, orderID int) error {
	orders, err := op.storage.GetOrders()
	if err != nil {
		return err
	}
	for _, order := range orders {
		if orderID != order.ID {
			continue
		}
		if clientID != order.ClientID {
			break
		}
		if !order.IsIssued {
			return ErrOrderIsNotIssued
		}
		if order.IsReturned {
			return ErrOrderIsReturned
		}
		maxReturnTime := order.OrderIssueDate.Add(time.Hour * 24 * 2)
		if maxReturnTime.Before(time.Now()) {
			return ErrReturnTimeExpired
		}
		return op.storage.ReturnOrderStorage(clientID, orderID)
	}
	return ErrClientOrderNotFound
}
