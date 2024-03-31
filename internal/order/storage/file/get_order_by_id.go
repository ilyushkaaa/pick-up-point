package storage

import (
	"context"

	"homework/internal/order/model"
	"homework/internal/order/storage"
)

func (fs *FileOrderStorage) GetOrderByID(ctx context.Context, ID uint64) (*model.Order, error) {
	orders, err := fs.GetOrders(ctx)
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		if order.ID == ID {
			return &order, nil
		}
	}
	return nil, storage.ErrOrderNotFound
}
