package delivery

import (
	"go.uber.org/zap"
	"homework/internal/cache"
	"homework/internal/pick-up_point/service"
)

type PPDelivery struct {
	cache   cache.Cache
	logger  *zap.SugaredLogger
	service service.PickUpPointService
}

func New(cache cache.Cache, service service.PickUpPointService, logger *zap.SugaredLogger) *PPDelivery {
	return &PPDelivery{
		cache:   cache,
		logger:  logger,
		service: service,
	}
}
