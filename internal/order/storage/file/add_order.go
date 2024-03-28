package storage

import (
	"context"

	"homework/internal/order/model"
)

func (fs *FileOrderStorage) AddOrder(ctx context.Context, newOrder model.Order) error {
	orders, err := fs.GetOrders(ctx)
	if err != nil {
		return err
	}
	orders = append(orders, newOrder)
	return fs.writeOrders(orders)
}
