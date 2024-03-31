package service

import (
	"context"
	"time"
)

func (op *OrderServicePP) DeleteOrder(ctx context.Context, orderID uint64) error {
	order, err := op.storage.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}
	if order.IsReturned {
		return op.storage.DeleteOrder(ctx, orderID)
	}
	if order.IsIssued {
		return ErrOrderAlreadyIssued
	}
	if order.StorageExpirationDate.After(time.Now()) {
		return ErrOrderShelfLifeNotExpired
	}
	return op.storage.DeleteOrder(ctx, orderID)

}
