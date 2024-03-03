package service

import (
	"homework/Homework-1/internal/order/myerrors"
)

func (op *OrderServicePP) IssueOrderService(orderIDs map[int]struct{}) error {
	orders, err := op.storage.GetOrders()
	if err != nil {
		return err
	}
	ordersCount := 0
	clientID := 0
	clientIDWasSet := false
	for _, ord := range orders {
		if _, exists := orderIDs[ord.ID]; !exists {
			continue
		}
		if clientIDWasSet && clientID != ord.ClientID {
			return myerrors.ErrorOrdersOfDifferentClients
		}
		if ord.IsIssued {
			return myerrors.ErrorOrderAlreadyIssued
		}
		clientIDWasSet = true
		clientID = ord.ClientID
		ordersCount++
	}
	if len(orderIDs) != ordersCount {
		return myerrors.ErrorOrderNotFound
	}
	return op.storage.IssueOrdersStorage(orderIDs)

}
