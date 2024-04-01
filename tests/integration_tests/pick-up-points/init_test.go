package pick_up_points

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	delivery "homework/internal/pick-up_point/delivery/http"
	"homework/internal/pick-up_point/service"
	storage "homework/internal/pick-up_point/storage/database"
	database "homework/pkg/database/postgres"
)

func initTest(t *testing.T) (*delivery.PPDelivery, database.DBops) {
	t.Helper()

	err := godotenv.Load("../../../.env")
	assert.NoError(t, err)
	db, err := database.NewDB(context.Background(), database.OptionTest)

	assert.NoError(t, err)

	st := storage.New(db)
	srv := service.New(st)
	dev := delivery.New(srv, zap.NewNop().Sugar())
	return dev, db
}
