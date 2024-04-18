package service

import (
	"context"
	"fmt"
	"time"

	"homework/internal/order/model"
)

func (op *OrderServicePP) DeleteOrder(ctx context.Context, orderID uint64) error {
	return op.transactionManager.RunRepeatableRead(ctx,
		func(ctxTX context.Context) error {
			gotFromCache := false
			var order *model.Order
			data, err := op.cache.GetFromCache(ctx, fmt.Sprintf("order_%d", orderID))
			if err == nil {
				order, gotFromCache = data.(*model.Order)
			}
			if !gotFromCache {
				order, err = op.orderStorage.GetOrderByID(ctx, orderID)
				if err != nil {
					return err
				}
			}
			if order.IsReturned {
				op.cache.GoDeleteFromCache(context.Background(), fmt.Sprintf("order_%d", orderID))
				return op.orderStorage.DeleteOrder(ctx, orderID)
			}
			if order.IsIssued {
				return ErrOrderAlreadyIssued
			}
			if order.StorageExpirationDate.After(time.Now()) {
				return ErrOrderShelfLifeNotExpired
			}
			err = op.orderStorage.DeleteOrder(ctx, orderID)
			if err != nil {
				op.cache.GoDeleteFromCache(context.Background(), fmt.Sprintf("order_%d", orderID))
			}
			return err
		})

}
