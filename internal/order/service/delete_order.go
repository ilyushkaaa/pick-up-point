package service

import (
	"context"
	"strconv"
	"time"

	"homework/internal/order/model"
)

func (op *OrderServicePP) DeleteOrder(ctx context.Context, orderID uint64) error {
	return op.transactionManager.RunRepeatableRead(ctx,
		func(ctxTX context.Context) error {
			gotFromCache := false
			var order *model.Order
			data, err := op.cacheOrderByID.GetFromCache(ctx, strconv.FormatUint(order.ID, 10))
			if err == nil {
				order, gotFromCache = data.(*model.Order)
			}
			if !gotFromCache {
				order, err = op.orderStorage.GetOrderByID(ctx, orderID)
				if err != nil {
					return err
				}
				op.cacheOrderByID.GoAddToCache(context.Background(), strconv.FormatUint(order.ID, 10), order)
			}
			if order.IsReturned {
				err = op.orderStorage.DeleteOrder(ctx, orderID)
				if err != nil {
					op.cacheOrderByID.GoDeleteFromCache(context.Background(), strconv.FormatUint(order.ID, 10))
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
				op.cacheOrderByID.GoDeleteFromCache(context.Background(), strconv.FormatUint(order.ID, 10))
			}
			return err
		})

}
