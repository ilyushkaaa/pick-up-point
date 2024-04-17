package service

import (
	"context"

	filtermodel "homework/internal/filters/model"
	ordermodel "homework/internal/order/model"
)

func (op *OrderServicePP) GetUserOrders(ctx context.Context, clientID uint64, filters filtermodel.Filters) ([]ordermodel.Order, error) {
	orders, err := op.orderStorage.GetUserOrders(ctx, clientID)
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
