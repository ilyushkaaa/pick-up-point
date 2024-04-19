//go:build integration
// +build integration

package pick_up_points

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	cacheInMemory "homework/internal/cache/in_memory"
	cacheRedis "homework/internal/cache/redis"
	storageOrder "homework/internal/order/storage/database"
	delivery "homework/internal/pick-up_point/delivery/http"
	"homework/internal/pick-up_point/service"
	storagePP "homework/internal/pick-up_point/storage/database"
	"homework/pkg/infrastructure/database/postgres"
	"homework/pkg/infrastructure/database/postgres/transaction_manager"
)

func initTest(t *testing.T) (*delivery.PPDelivery, database.Database) {
	t.Helper()

	tm, err := transaction_manager.New(context.Background())

	require.NoError(t, err)

	db := database.New(tm)

	stPP := storagePP.New(db)
	stOrder := storageOrder.New(db)
	logger := zap.NewNop().Sugar()
	redisCache := cacheRedis.New(logger)
	t.Cleanup(func() {
		err = redisCache.Close()
		assert.NoError(t, err)
	})
	imMemoryCache := cacheInMemory.New(logger)
	srv := service.New(stPP, stOrder, tm, imMemoryCache)
	dev := delivery.New(redisCache, srv, logger)
	return dev, db
}
