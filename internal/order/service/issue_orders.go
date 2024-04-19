package service

import (
	"context"
	"fmt"

	cache "homework/internal/cache/in_memory"
)

func (op *OrderServicePP) IssueOrders(ctx context.Context, orderIDs []uint64) error {
	orderIDsMap := make(map[uint64]struct{}, len(orderIDs))
	for _, id := range orderIDs {
		orderIDsMap[id] = struct{}{}
	}
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
				if _, exists := orderIDsMap[order.ID]; !exists {
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
			err = op.orderStorage.IssueOrders(ctx, orderIDs)
			if err != nil {
				op.cache.GoDeleteFromCache(context.Background(), getKeysOrderKeys(orderIDs)...)
			}
			return err
		})
}

func getKeysOrderKeys(IDs []uint64) []string {
	keys := make([]string, 0, len(IDs))
	for _, id := range IDs {
		keys = append(keys, fmt.Sprintf("%s_%d", cache.PrefixOrderByID, id))
	}
	return keys
}
