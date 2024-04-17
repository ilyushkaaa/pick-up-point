package service

import (
	"context"

	"homework/internal/order/model"
)

func (op *OrderServicePP) GetOrderReturns(ctx context.Context, maxOrdersPerPage, pageNum uint64) ([]model.Order, error) {
	orders, err := op.orderStorage.GetOrderReturns(ctx)
	if err != nil {
		return nil, err
	}
	startingOrderForPage := maxOrdersPerPage * (pageNum - 1)
	if startingOrderForPage >= uint64(len(orders)) {
		return nil, ErrNoOrdersOnThisPage
	}
	return orders[startingOrderForPage:min(startingOrderForPage+4, uint64(len(orders)))], nil
}
