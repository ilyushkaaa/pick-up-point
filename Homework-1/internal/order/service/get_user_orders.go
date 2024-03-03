package service

import "homework/Homework-1/internal/order/model"

func (op *OrderServicePP) GetUserOrdersService(clientID, limit int, ppOnly bool) ([]model.Order, error) {
	orders, err := op.storage.GetUserOrdersStorage(clientID)
	if err != nil {
		return nil, err
	}
	if limit == 0 {
		limit = len(orders)
	}
	ordersToReturn := make([]model.Order, 0)
	for _, ord := range orders {
		if ppOnly {
			if !ord.IsIssued {
				ordersToReturn = append(ordersToReturn, ord)
			} else {
				continue
			}
		} else {
			ordersToReturn = append(ordersToReturn, ord)
		}
		limit--
		if limit == 0 {
			break
		}
	}
	return ordersToReturn, nil
}
