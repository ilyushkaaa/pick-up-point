package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"homework/internal/order/model"
	"homework/internal/order/storage"
)

func (op *OrderServicePP) AddOrder(ctx context.Context, order model.Order) error {
	_, err := op.cacheOrderByID.GetFromCache(ctx, strconv.FormatUint(order.ID, 10))
	if err == nil {
		return ErrOrderAlreadyInPickUpPoint
	}
	orderByID, err := op.orderStorage.GetOrderByID(ctx, order.ID)
	if err != nil && !errors.Is(err, storage.ErrOrderNotFound) {
		return err
	}
	if err == nil {
		op.cacheOrderByID.GoAddToCache(context.Background(), strconv.FormatUint(order.ID, 10), orderByID)
		return ErrOrderAlreadyInPickUpPoint
	}
	return op.transactionManager.RunRepeatableRead(ctx,
		func(ctxTX context.Context) error {
			gotFromCache := false
			_, err = op.cachePPByID.GetFromCache(ctx, strconv.FormatUint(order.PickUpPointID, 10))
			if err == nil {
				gotFromCache = true
			}
			if !gotFromCache {
				_, err = op.ppStorage.GetPickUpPointByID(ctx, order.PickUpPointID)
				if err != nil {
					return err
				}
				op.cachePPByID.GoAddToCache(context.Background(), strconv.FormatUint(order.PickUpPointID, 10), order)
			}
			if order.PackageType != "" {
				chosenPackage, exists := op.packages[order.PackageType]
				if !exists {
					return ErrUnknownPackage
				}
				if chosenPackage.MaxWeight != 0 && chosenPackage.MaxWeight < order.Weight {
					return ErrPackageCanNotBeApplied
				}
				order.Price += chosenPackage.Price
			}
			if time.Now().After(order.StorageExpirationDate) {
				return ErrShelfTimeExpired
			}
			err = op.orderStorage.AddOrder(ctx, order)
			if err == nil {
				op.metrics.DeliveredOrdersCount.Inc()
				op.metrics.OrdersCountByPickUpPoints.WithLabelValues(strconv.FormatUint(order.PickUpPointID, 10)).Inc()
				op.metrics.OrdersByClientCount.WithLabelValues(strconv.FormatUint(order.ClientID, 10)).Inc()
			}
			return err

		})
}
