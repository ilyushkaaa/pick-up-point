package service

import (
	"context"
	"time"
)

func (op *OrderServicePP) DeleteOrder(ctx context.Context, orderID uint64) error {
	return op.transactionManager.RunRepeatableRead(ctx,
		func(ctxTX context.Context) error {
			order, err := op.orderStorage.GetOrderByID(ctx, orderID)
			if err != nil {
				return err
			}
			if order.IsReturned {
				return op.orderStorage.DeleteOrder(ctx, orderID)
			}
			if order.IsIssued {
				return ErrOrderAlreadyIssued
			}
			if order.StorageExpirationDate.After(time.Now()) {
				return ErrOrderShelfLifeNotExpired
			}
			return op.orderStorage.DeleteOrder(ctx, orderID)
		})

}
