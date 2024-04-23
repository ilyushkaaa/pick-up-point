package pick_up_points

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/tests/test_pb"
)

func TestUpdatePickUpPoint(t *testing.T) {

	t.Run("error pick-up point with such id was not found", func(t *testing.T) {
		del := setUp(t, tableName)

		result, err := del.Update(context.Background(), &test_pb.UpdatePPNotExist)

		assert.ErrorIs(t, err, status.Error(codes.NotFound, "no pick-up points with such id"))
		assert.Nil(t, result)
	})

	t.Run("error pick-up point with such name already exists", func(t *testing.T) {
		del := setUp(t, tableName)

		result, err := del.Update(context.Background(), &test_pb.UpdatePPNameAlreadyExists)

		assert.ErrorIs(t, err, status.Error(codes.InvalidArgument, "pick-up point with such name already exists"))
		assert.Nil(t, result)
	})

	t.Run("ok", func(t *testing.T) {
		del := setUp(t, tableName)

		result, err := del.Update(context.Background(), &test_pb.UpdatePPOk)

		assert.NoError(t, err)
		assert.Equal(t, result, &test_pb.UpdatePPOk)
	})
}
