package pick_up_points

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/tests/test_pb"
)

const tableName = "pick_up_points"

func TestAddPickUpPoint(t *testing.T) {

	t.Run("error pick-up point with such name already exists", func(t *testing.T) {
		del := setUp(t, tableName)

		result, err := del.Add(context.Background(), &test_pb.ValidPPAddRequest)

		assert.ErrorIs(t, err, status.Error(codes.InvalidArgument, "pick-up point with such name already exists"))
		assert.Nil(t, result)
	})

	t.Run("ok", func(t *testing.T) {
		del := setUp(t, tableName)

		result, err := del.Add(context.Background(), &test_pb.ValidPPAddRequestUnique)

		assert.NoError(t, err)
		assert.Equal(t, result, &test_pb.AddedPP)
	})
}
