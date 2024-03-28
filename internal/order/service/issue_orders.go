package service

import "context"

func (op *OrderServicePP) IssueOrders(ctx context.Context, orderIDs map[uint64]struct{}) error {
	orders, err := op.storage.GetOrders(ctx)
	if err != nil {
		return err
	}
	ordersCount := 0
	clientID := uint64(0)
	clientIDWasSet := false
	for _, order := range orders {
		if _, exists := orderIDs[order.ID]; !exists {
			continue
		}
		if clientIDWasSet && clientID != order.ClientID {
			return ErrOrdersOfDifferentClients
		}
		if order.IsIssued {
			return ErrOrderAlreadyIssued
		}
		clientIDWasSet = true
		clientID = order.ClientID
		ordersCount++
	}
	if len(orderIDs) != ordersCount {
		return ErrNotAllOrdersWereFound
	}
	return op.storage.IssueOrders(ctx, orderIDs)

}
