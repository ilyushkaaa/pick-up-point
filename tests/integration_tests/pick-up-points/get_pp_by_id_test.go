package pick_up_points

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/tests/test_pb"
)

func TestGetPickUpPointByID(t *testing.T) {

	t.Run("error no pick-up points with such id", func(t *testing.T) {
		del := setUp(t, tableName)

		result, err := del.GetByID(context.Background(), &test_pb.GetByIDRequestNotExist)

		assert.ErrorIs(t, err, status.Error(codes.NotFound, "no pick-up points with such id"))
		assert.Nil(t, result)
	})

	t.Run("ok", func(t *testing.T) {
		del := setUp(t, tableName)

		result, err := del.GetByID(context.Background(), &test_pb.GetByIDRequestOk)

		assert.NoError(t, err)
		assert.Equal(t, result, &test_pb.GetPPByIDSuccess)
	})
}
