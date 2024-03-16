package delivery

import "homework/internal/pick-up_point/service"

type PPDelivery struct {
	service service.PickUpPointService
}

func New(service service.PickUpPointService) *PPDelivery {
	return &PPDelivery{
		service: service,
	}
}
