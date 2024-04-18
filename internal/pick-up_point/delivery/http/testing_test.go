package delivery

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	cacheRedis "homework/internal/cache/redis"
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
	logger := zap.NewNop().Sugar()
	redisCache := cacheRedis.New(logger)
	t.Cleanup(func() {
		err := redisCache.Close()
		assert.NoError(t, err)
	})
	del := PPDelivery{redisCache, logger, mockPPService}
	return pickUpPointDeliveryFixtures{
		ctrl:        ctrl,
		del:         del,
		mockService: mockPPService,
	}
}
