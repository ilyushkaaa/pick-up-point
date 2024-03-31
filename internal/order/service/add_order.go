package service

import (
	"context"
	"errors"
	"time"

	"homework/internal/order/model"
	"homework/internal/order/storage"
)

func (op *OrderServicePP) AddOrder(ctx context.Context, order model.Order) error {
	_, err := op.storage.GetOrderByID(ctx, order.ID)
	if err != nil && !errors.Is(err, storage.ErrOrderNotFound) {
		return err
	}
	if err == nil {
		return ErrOrderAlreadyInPickUpPoint
	}
	if order.PackageType != "" {
		chosenPackage, exists := op.packages[order.PackageType]
		if !exists {
			return ErrUnknownPackage
		}
		if ok := chosenPackage.CheckIfPackageCanBeApplied(order); !ok {
			return ErrPackageCanNotBeApplied
		}
		order.Price += chosenPackage.GetPrice()
	}
	if time.Now().After(order.StorageExpirationDate) {
		return ErrShelfTimeExpired
	}
	return op.storage.AddOrder(ctx, order)
}
