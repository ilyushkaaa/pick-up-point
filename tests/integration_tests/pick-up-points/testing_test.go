//go:build integration
// +build integration

package pick_up_points

import (
	"context"
	"fmt"
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
	"homework/tests/fixtures"
	"homework/tests/states"
)

func setUp(t *testing.T, tableName string) *delivery.PPDelivery {
	t.Helper()
	ctx := context.Background()

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
	del := delivery.New(redisCache, srv, logger)

	err = truncateTable(ctx, db, tableName)
	require.NoError(t, err)

	fillDataBase(t, db)
	return del
}

func truncateTable(ctx context.Context, db database.Database, tableName string) error {
	_, err := db.Exec(ctx, fmt.Sprintf("TRUNCATE table %s RESTART IDENTITY", tableName))
	return err
}

func fillDataBase(t *testing.T, db database.Database) {
	t.Helper()

	ctx := context.Background()
	pp := fixtures.PickUpPoint().Valid().V()
	_, err := db.Exec(ctx,
		`INSERT INTO pick_up_points (id, name, phone_number, region, city, street, house_num)
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		pp.ID, pp.Name, pp.PhoneNumber, pp.Address.Region, pp.Address.City, pp.Address.Street, pp.Address.HouseNum)
	require.NoError(t, err)

	pp = fixtures.PickUpPoint().Valid().ID(states.PPID2).Name(states.PPName2).V()
	_, err = db.Exec(ctx,
		`INSERT INTO pick_up_points (id, name, phone_number, region, city, street, house_num)
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		pp.ID, pp.Name, pp.PhoneNumber, pp.Address.Region, pp.Address.City, pp.Address.Street, pp.Address.HouseNum)
	require.NoError(t, err)
}
