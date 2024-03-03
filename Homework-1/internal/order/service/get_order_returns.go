package service

import "homework/Homework-1/internal/order/model"

const maxOrdersPerPage = 4

func (op *OrderServicePP) GetOrderReturnsService() ([][]model.Order, error) {
	orders, err := op.storage.GetOrderReturnsStorage()
	if err != nil {
		return nil, err
	}
	orderPages := make([][]model.Order, 0)
	for len(orders) > maxOrdersPerPage {
		orderPages = append(orderPages, orders[:maxOrdersPerPage])
		orders = orders[maxOrdersPerPage:]
	}
	orderPages = append(orderPages, orders)
	return orderPages, nil
}
