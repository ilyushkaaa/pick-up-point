package service

import (
	filtermodel "homework/internal/filters/model"
	ordermodel "homework/internal/order/model"
)

func (op *OrderServicePP) GetUserOrdersService(clientID int, filters filtermodel.Filters) ([]ordermodel.Order, error) {
	orders, err := op.storage.GetUserOrdersStorage(clientID)
	if err != nil {
		return nil, err
	}
	if filters.Limit == 0 {
		filters.Limit = len(orders)
	}
	ordersToReturn := make([]ordermodel.Order, 0)
	for _, order := range orders {
		if filters.PPOnly {
			if !order.IsIssued {
				ordersToReturn = append(ordersToReturn, order)
			} else {
				continue
			}
		} else {
			ordersToReturn = append(ordersToReturn, order)
		}
		filters.Limit--
		if filters.Limit == 0 {
			break
		}
	}
	return ordersToReturn, nil
}
