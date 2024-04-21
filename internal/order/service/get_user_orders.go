package service

import (
	"context"
	"strconv"

	filtermodel "homework/internal/filters/model"
	ordermodel "homework/internal/order/model"
)

func (op *OrderServicePP) GetUserOrders(ctx context.Context, clientID uint64, filters filtermodel.Filters) ([]ordermodel.Order, error) {
	gotFromCache := false
	var orders []ordermodel.Order
	data, err := op.cacheOrdersByClient.GetFromCache(ctx, strconv.FormatUint(clientID, 10))
	if err == nil {
		orders, gotFromCache = data.([]ordermodel.Order)
	}
	if !gotFromCache {
		orders, err = op.orderStorage.GetUserOrders(ctx, clientID)
		if err != nil {
			return nil, err
		}
		op.cacheOrdersByClient.GoAddToCache(context.Background(), strconv.FormatUint(clientID, 10), orders)
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
	op.cacheOrdersByClient.GoAddToCache(context.Background(), strconv.FormatUint(clientID, 10), ordersToReturn)
	return ordersToReturn, nil
}
