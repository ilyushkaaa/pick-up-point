package grpc

import (
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"homework/internal/order/service"
	pb "homework/internal/pb/order"
)

type OrderDelivery struct {
	service service.OrderService
	logger  *zap.SugaredLogger
	tracer  trace.Tracer

	pb.UnimplementedOrdersServer
}

func New(service service.OrderService, logger *zap.SugaredLogger, tracer trace.Tracer) *OrderDelivery {
	return &OrderDelivery{
		UnimplementedOrdersServer: pb.UnimplementedOrdersServer{},
		service:                   service,
		logger:                    logger,
		tracer:                    tracer,
	}
}
