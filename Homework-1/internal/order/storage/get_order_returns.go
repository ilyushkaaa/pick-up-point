package storage

import (
	"homework/Homework-1/internal/order/model"
)

func (fs *FileOrderStorage) GetOrderReturnsStorage() ([]model.Order, error) {
	orders, err := fs.GetOrders()
	if err != nil {
		return nil, err
	}
	returnedOrders := make([]model.Order, 0)
	for _, ord := range orders {
		if ord.IsReturned {
			returnedOrders = append(returnedOrders, ord)
		}
	}
	return returnedOrders, nil
}
