package service

import "homework/internal/order/model"

const maxOrdersPerPage = 4

func (op *OrderServicePP) GetOrderReturnsService(pageNum int) ([]model.Order, error) {
	orders, err := op.storage.GetOrderReturnsStorage()
	if err != nil {
		return nil, err
	}
	startingOrderForPage := maxOrdersPerPage * (pageNum - 1)
	if startingOrderForPage >= len(orders) {
		return nil, ErrNoOrdersOnThisPage
	}
	return orders[startingOrderForPage:min(startingOrderForPage+4, len(orders))], nil
}
