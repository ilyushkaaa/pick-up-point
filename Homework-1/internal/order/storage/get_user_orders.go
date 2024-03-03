package storage

import "homework/Homework-1/internal/order/model"

func (fs *FileOrderStorage) GetUserOrdersStorage(clientID int) ([]model.Order, error) {
	orders, err := fs.GetOrders()
	if err != nil {
		return nil, err
	}
	ordersToReturn := make([]model.Order, 0)
	for _, ord := range orders {
		if ord.ClientID == clientID {
			ordersToReturn = append(ordersToReturn, ord)
		}
	}
	return ordersToReturn, nil
}
