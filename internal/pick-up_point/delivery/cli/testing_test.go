package delivery

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_service "homework/internal/pick-up_point/service/mocks"
)

type pickUpPointDeliveryFixtures struct {
	ctrl        *gomock.Controller
	del         PPDelivery
	mockService *mock_service.MockPickUpPointService
}

func setUp(t *testing.T) pickUpPointDeliveryFixtures {
	ctrl := gomock.NewController(t)
	mockPPService := mock_service.NewMockPickUpPointService(ctrl)
	del := PPDelivery{mockPPService}
	return pickUpPointDeliveryFixtures{
		ctrl:        ctrl,
		del:         del,
		mockService: mockPPService,
	}
}
