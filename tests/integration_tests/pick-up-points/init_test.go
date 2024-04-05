package pick_up_points

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	delivery "homework/internal/pick-up_point/delivery/http"
	"homework/internal/pick-up_point/service"
	storage "homework/internal/pick-up_point/storage/database"
	database "homework/pkg/database/postgres"
)

func initTest(t *testing.T) (*delivery.PPDelivery, database.Database) {
	t.Helper()

	db, err := database.New(context.Background())

	require.NoError(t, err)

	st := storage.New(db)
	srv := service.New(st)
	dev := delivery.New(srv, zap.NewNop().Sugar())
	return dev, db
}
