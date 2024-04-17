//go:build integration
// +build integration

package pick_up_points

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	storageOrder "homework/internal/order/storage/database"
	delivery "homework/internal/pick-up_point/delivery/http"
	"homework/internal/pick-up_point/service"
	storagePP "homework/internal/pick-up_point/storage/database"
	"homework/pkg/database/postgres"
)

func initTest(t *testing.T) (*delivery.PPDelivery, database.Database) {
	t.Helper()

	db, err := database.New(context.Background())

	require.NoError(t, err)

	stPP := storagePP.New(db)
	stOrder := storageOrder.New(db)
	srv := service.New(stPP, stOrder)
	dev := delivery.New(srv, zap.NewNop().Sugar())
	return dev, db
}
