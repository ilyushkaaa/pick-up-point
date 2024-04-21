package service

import (
	"context"
	"strconv"
	"time"

	"homework/internal/order/model"
)

func (op *OrderServicePP) ReturnOrder(ctx context.Context, clientID, orderID uint64) error {
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
			if err == nil {
				op.cacheOrderByID.GoDeleteFromCache(context.Background(), strconv.FormatUint(order.ID, 10))
			}
			return err
		})

}
