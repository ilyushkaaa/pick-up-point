package delivery

import (
	"go.uber.org/zap"
	"homework/internal/pick-up_point/service"
)

type PPDelivery struct {
	logger  *zap.SugaredLogger
	service service.PickUpPointService
}

func New(service service.PickUpPointService, logger *zap.SugaredLogger) *PPDelivery {
	return &PPDelivery{
		logger:  logger,
		service: service,
	}
}
