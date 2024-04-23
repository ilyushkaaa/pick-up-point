package grpc

import (
	"go.uber.org/zap"
	"homework/internal/order/service"
	pb "homework/internal/pb/order"
)

type OrderDelivery struct {
	service service.OrderService
	logger  *zap.SugaredLogger

	pb.UnimplementedOrdersServer
}

func New(service service.OrderService, logger *zap.SugaredLogger) *OrderDelivery {
	return &OrderDelivery{
		UnimplementedOrdersServer: pb.UnimplementedOrdersServer{},
		service:                   service,
		logger:                    logger,
	}
}
