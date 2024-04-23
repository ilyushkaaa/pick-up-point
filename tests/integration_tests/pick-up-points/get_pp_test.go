//go:build integration
// +build integration

package pick_up_points

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	pb "homework/internal/pb/pick-up_point"
	"homework/tests/test_pb"
)

func TestGetPickUpPoints(t *testing.T) {

	t.Run("ok", func(t *testing.T) {
		del := setUp(t, tableName)
		fakeStream := &fakeServerStream{
			pickUpPoints: make([]*pb.PickUpPoint, 0),
		}

		err := del.GetAll(&pb.GetAllRequest{}, fakeStream)

		assert.NoError(t, err)
		assert.Equal(t, fakeStream.pickUpPoints, test_pb.GetAllPP)
	})
}

type fakeServerStream struct {
	pickUpPoints []*pb.PickUpPoint
	grpc.ServerStream
}

func (f *fakeServerStream) Send(data *pb.PickUpPoint) error {
	f.pickUpPoints = append(f.pickUpPoints, data)
	return nil
}

func (f *fakeServerStream) Context() context.Context {
	return context.Background()
}
