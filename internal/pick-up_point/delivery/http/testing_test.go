package delivery

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	ttl, err := strconv.ParseUint(os.Getenv("CACHE_REDIS_TTL"), 10, 64)
	require.NoError(t, err)
	redisCache := cacheRedis.New(logger, fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASSWORD"), time.Duration(ttl))
	t.Cleanup(func() {
		err = redisCache.Close()
		assert.NoError(t, err)
	})
	del := PPDelivery{redisCache, logger, mockPPService}
	return pickUpPointDeliveryFixtures{
		ctrl:        ctrl,
		del:         del,
		mockService: mockPPService,
	}
}
