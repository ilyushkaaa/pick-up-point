//go:build integration
// +build integration

package events

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"homework/pkg/infrastructure/kafka/consumer"
)

const testTimeout = time.Second * 7

func TestLoggingEvents(t *testing.T) {

	t.Run("test get", func(t *testing.T) {
		s := setUpAndConsume(t)
		defer s.cancel()
		var addr address = "127.0.0.1:12345"
		method := "/pb.Orders/IssueOrders"
		ctx := peer.NewContext(context.Background(), &peer.Peer{Addr: addr})
		info := &grpc.UnaryServerInfo{
			FullMethod: method,
		}
		require.NoError(t, consumer.WaitForConsumerReady(s.waitChan))

		s.i.AccessLog(ctx, "", info, func(ctx context.Context, req any) (any, error) {
			return nil, nil
		})

		select {
		case <-s.waitChan:
			assert.Contains(t, s.buf.String(), `"method": "/pb.Orders/IssueOrders"`)
			assert.Contains(t, s.buf.String(), `"remote_addr": "127.0.0.1:12345"`)
		case <-time.After(testTimeout):
			t.Error("Timeout occurred")
		}

	})

	t.Run("test post", func(t *testing.T) {
		s := setUpAndConsume(t)
		defer s.cancel()
		var addr address = "127.0.0.1:34567"
		method := "/pb.PicUpPoints/Add"
		ctx := peer.NewContext(context.Background(), &peer.Peer{Addr: addr})
		info := &grpc.UnaryServerInfo{
			FullMethod: method,
		}
		require.NoError(t, consumer.WaitForConsumerReady(s.waitChan))

		s.i.AccessLog(ctx, "", info, func(ctx context.Context, req any) (any, error) {
			return nil, nil
		})

		select {
		case <-s.waitChan:
			assert.Contains(t, s.buf.String(), `"method": "/pb.PicUpPoints/Add"`)
			assert.Contains(t, s.buf.String(), `"remote_addr": "127.0.0.1:34567"`)
		case <-time.After(testTimeout):
			t.Error("Timeout occurred")
		}
	})

}
