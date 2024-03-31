package storage

import (
	"context"

	"homework/internal/order/model"
)

func (fs *FileOrderStorage) GetOrderReturns(ctx context.Context) ([]model.Order, error) {
	orders, err := fs.GetOrders(ctx)
	if err != nil {
		return nil, err
	}
	returnedOrders := make([]model.Order, 0)
	for _, order := range orders {
		if order.IsReturned {
			returnedOrders = append(returnedOrders, order)
		}
	}
	return returnedOrders, nil
}
