package storage

import (
	"context"

	"homework/internal/order/model"
)

func (fs *FileOrderStorage) GetUserOrders(ctx context.Context, clientID uint64) ([]model.Order, error) {
	orders, err := fs.GetOrders(ctx)
	if err != nil {
		return nil, err
	}
	ordersToReturn := make([]model.Order, 0)
	for _, order := range orders {
		if order.ClientID == clientID {
			ordersToReturn = append(ordersToReturn, order)
		}
	}
	return ordersToReturn, nil
}
