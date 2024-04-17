package service

import (
	"context"
)

func (op *OrderServicePP) IssueOrders(ctx context.Context, orderIDs map[uint64]struct{}) error {
	return op.transactionManager.RunRepeatableRead(ctx,
		func(ctxTX context.Context) error {
			orders, err := op.orderStorage.GetOrders(ctx)
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
			return op.orderStorage.IssueOrders(ctx, orderIDs)
		})
}
