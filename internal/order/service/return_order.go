package service

import (
	"context"
	"time"
)

func (op *OrderServicePP) ReturnOrder(ctx context.Context, clientID, orderID uint64) error {
	return op.transactionManager.RunRepeatableRead(ctx,
		func(ctxTX context.Context) error {
			order, err := op.orderStorage.GetOrderByID(ctx, orderID)
			if err != nil {
				return err
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
			return op.orderStorage.ReturnOrder(ctx, clientID, orderID)
		})

}
