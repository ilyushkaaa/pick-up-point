package service

import (
	"context"
	"errors"
	"time"

	"homework/internal/order/model"
	"homework/internal/order/storage"
)

func (op *OrderServicePP) AddOrder(ctx context.Context, order model.Order) error {
	_, err := op.orderStorage.GetOrderByID(ctx, order.ID)
	if err != nil && !errors.Is(err, storage.ErrOrderNotFound) {
		return err
	}
	if err == nil {
		return ErrOrderAlreadyInPickUpPoint
	}
	return op.transactionManager.RunRepeatableRead(ctx,
		func(ctxTX context.Context) error {
			_, err = op.ppStorage.GetPickUpPointByID(ctx, order.PickUpPointID)
			if err != nil {
				return err
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
			return op.orderStorage.AddOrder(ctx, order)

		})
}
