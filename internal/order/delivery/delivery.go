package delivery

import (
	"homework/internal/order/service"
)

type OrderDelivery struct {
	service service.OrderService
}

func New(service service.OrderService) *OrderDelivery {
	return &OrderDelivery{
		service: service,
	}
}
