package service

import (
	"context"
	"errors"
	"time"

	"homework/internal/order/model"
	"homework/internal/order/storage"
)

func (op *OrderServicePP) AddOrder(ctx context.Context, order model.Order, packageType string) error {
	_, err := op.storage.GetOrderByID(ctx, order.ID)
	if err != nil && !errors.Is(err, storage.ErrOrderNotFound) {
		return err
	}
	if err == nil {
		return ErrOrderAlreadyInPickUpPoint
	}
	if packageType != "" {
		chosenPackage, exists := op.packages[packageType]
		if !exists {
			return ErrUnknownPackage
		}
		err = chosenPackage.AddPackageToOrder(&order)
		if err != nil {
			return err
		}
	}
	if time.Now().After(order.StorageExpirationDate) {
		return ErrShelfTimeExpired
	}
	return op.storage.AddOrder(ctx, order)
}
