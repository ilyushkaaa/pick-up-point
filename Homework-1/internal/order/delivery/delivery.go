package delivery

import (
	"homework/Homework-1/internal/order/service"
)

type OrderDelivery struct {
	service service.OrderService
}

func NewOrderDelivery(service service.OrderService) *OrderDelivery {
	return &OrderDelivery{
		service: service,
	}
}
