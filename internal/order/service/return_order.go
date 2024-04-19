package service

import (
	"context"
	"fmt"
	"time"

	cache "homework/internal/cache/in_memory"
	"homework/internal/order/model"
)

func (op *OrderServicePP) ReturnOrder(ctx context.Context, clientID, orderID uint64) error {
	return op.transactionManager.RunRepeatableRead(ctx,
		func(ctxTX context.Context) error {
			gotFromCache := false
			var order *model.Order
			data, err := op.cache.GetFromCache(ctx, fmt.Sprintf("%s_%d", cache.PrefixOrderByID, order.ID))
			if err == nil {
				order, gotFromCache = data.(*model.Order)
			}
			if !gotFromCache {
				order, err = op.orderStorage.GetOrderByID(ctx, orderID)
				if err != nil {
					return err
				}
				op.cache.GoAddToCache(context.Background(), fmt.Sprintf("%s_%d", cache.PrefixOrderByID, order.ID), order)
			}
			if order.ClientID != clientID {
				return ErrClientOrderNotFound
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

			err = op.orderStorage.ReturnOrder(ctx, clientID, orderID)
			if err != nil {
				op.cache.GoDeleteFromCache(context.Background(), fmt.Sprintf("%s_%d", cache.PrefixOrderByID, order.ID))
			}
			return err
		})

}
