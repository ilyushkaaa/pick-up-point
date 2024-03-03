package storage

import (
	"homework/Homework-1/internal/order/model"
)

func (fs *FileOrderStorage) AddOrderStorage(newOrder model.Order) error {
	orders, err := fs.GetOrders()
	if err != nil {
		return err
	}
	orders = append(orders, newOrder)
	return fs.writeOrders(orders)
}
