package delivery

import (
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"homework/internal/cache"
	pb "homework/internal/pb/pick-up_point"
	"homework/internal/pick-up_point/service"
)

type PPDelivery struct {
	cache   cache.Cache
	logger  *zap.SugaredLogger
	service service.PickUpPointService
	tracer  trace.Tracer

	pb.UnimplementedPickUpPointsServer
}

func New(cache cache.Cache, service service.PickUpPointService, logger *zap.SugaredLogger, tracer trace.Tracer) *PPDelivery {
	return &PPDelivery{
		UnimplementedPickUpPointsServer: pb.UnimplementedPickUpPointsServer{},
		cache:                           cache,
		logger:                          logger,
		service:                         service,
		tracer:                          tracer,
	}
}
