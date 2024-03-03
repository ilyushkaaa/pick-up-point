package service

import (
	"time"

	"homework/Homework-1/internal/order/myerrors"
)

func (op *OrderServicePP) ReturnOrderService(clientID, orderID int) error {
	orders, err := op.storage.GetOrders()
	if err != nil {
		return err
	}
	for _, ord := range orders {
		if orderID != ord.ID {
			continue
		}
		if clientID == ord.ClientID {
			if !ord.IsIssued {
				return myerrors.ErrorOrderIsNotIssued
			}
			if ord.IsReturned {
				return myerrors.ErrorOrderIsReturned
			}
			maxReturnTime := ord.OrderIssueDate.Add(time.Hour * 24 * 2)
			if maxReturnTime.Before(time.Now()) {
				return myerrors.ErrorReturnTimeExpired
			}
			return op.storage.ReturnOrderStorage(clientID, orderID)
		}
		break

	}
	return myerrors.ErrorClientOrderNotFound
}
