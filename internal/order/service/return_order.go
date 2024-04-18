package service

import (
	"context"
	"fmt"
	"time"

	"homework/internal/order/model"
)

func (op *OrderServicePP) ReturnOrder(ctx context.Context, clientID, orderID uint64) error {
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
				op.cache.GoDeleteFromCache(context.Background(), fmt.Sprintf("order_%d", orderID))
			}
			return err
		})

}
