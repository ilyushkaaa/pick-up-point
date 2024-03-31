package delivery

import (
	"go.uber.org/zap"
	"homework/internal/order/service"
)

type OrderDelivery struct {
	service service.OrderService
	logger  *zap.SugaredLogger
}

func New(service service.OrderService, logger *zap.SugaredLogger) *OrderDelivery {
	return &OrderDelivery{
		service: service,
		logger:  logger,
	}
}
