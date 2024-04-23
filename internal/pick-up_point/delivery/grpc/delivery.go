package delivery

import (
	"go.uber.org/zap"
	"homework/internal/cache"
	pb "homework/internal/pb/pick-up_point"
	"homework/internal/pick-up_point/service"
)

type PPDelivery struct {
	cache   cache.Cache
	logger  *zap.SugaredLogger
	service service.PickUpPointService

	pb.UnimplementedPickUpPointsServer
}

func New(cache cache.Cache, service service.PickUpPointService, logger *zap.SugaredLogger) *PPDelivery {
	return &PPDelivery{
		UnimplementedPickUpPointsServer: pb.UnimplementedPickUpPointsServer{},
		cache:                           cache,
		logger:                          logger,
		service:                         service,
	}
}
