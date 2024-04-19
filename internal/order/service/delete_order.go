package service

import (
	"context"
	"fmt"
	"time"

	cache "homework/internal/cache/in_memory"
	"homework/internal/order/model"
)

func (op *OrderServicePP) DeleteOrder(ctx context.Context, orderID uint64) error {
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
			if order.IsReturned {
				err = op.orderStorage.DeleteOrder(ctx, orderID)
				if err != nil {
					op.cache.GoDeleteFromCache(context.Background(), fmt.Sprintf("%s_%d", cache.PrefixOrderByID, order.ID))
				}
				return err
			}
			if order.IsIssued {
				return ErrOrderAlreadyIssued
			}
			if order.StorageExpirationDate.After(time.Now()) {
				return ErrOrderShelfLifeNotExpired
			}
			err = op.orderStorage.DeleteOrder(ctx, orderID)
			if err != nil {
				op.cache.GoDeleteFromCache(context.Background(), fmt.Sprintf("%s_%d", cache.PrefixOrderByID, order.ID))
			}
			return err
		})

}
