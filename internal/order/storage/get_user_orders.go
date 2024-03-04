package storage

import "homework/internal/order/model"

func (fs *FileOrderStorage) GetUserOrdersStorage(clientID int) ([]model.Order, error) {
	orders, err := fs.GetOrders()
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
