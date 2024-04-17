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
	"homework/pkg/database/postgres/transaction_manager"
)

func initTest(t *testing.T) (*delivery.PPDelivery, database.Database) {
	t.Helper()

	tm, err := transaction_manager.New(context.Background())

	require.NoError(t, err)

	db := database.New(tm)

	stPP := storagePP.New(db)
	stOrder := storageOrder.New(db)
	srv := service.New(stPP, stOrder, tm)
	dev := delivery.New(srv, zap.NewNop().Sugar())
	return dev, db
}
