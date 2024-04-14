package events

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	eventsProducer "homework/internal/events/service/producer"
	"homework/internal/middleware"
	"homework/pkg/kafka"
	"homework/pkg/kafka/consumer"
	"homework/pkg/kafka/producer"
)

const eventsTopic = "test_events"

type TestEventsFixtures struct {
	mw           *middleware.Middleware
	buf          *bytes.Buffer
	syncProducer *producer.SyncProducer
	brokers      []string
	logger       *zap.SugaredLogger
	cancelFunc   context.CancelFunc
}

type fakeHandler struct{}

func (h *fakeHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func setUpAndConsume(t *testing.T) TestEventsFixtures {
	t.Helper()

	var buf bytes.Buffer
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			MessageKey: "message",
		}),
		zapcore.AddSync(&buf),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)
	logger := zap.New(core)

	zapL := logger.Sugar()

	brokers := []string{"127.0.0.1:8004", "127.0.0.1:8005", "127.0.0.1:8006"}
	syncProducer, err := producer.New(brokers)
	require.NoError(t, err)

	t.Cleanup(func() {
		err = syncProducer.Close()
		require.NoError(t, err)
	})

	ep := eventsProducer.New(syncProducer, eventsTopic)
	mw := middleware.New(zapL, ep)

	ctx, cancel := context.WithCancel(context.Background())

	testFixtures := TestEventsFixtures{
		mw:           mw,
		buf:          &buf,
		syncProducer: syncProducer,
		brokers:      brokers,
		logger:       zapL,
		cancelFunc:   cancel,
	}
	testFixtures.GoRunConsume(t, ctx)

	return testFixtures
}

func (s TestEventsFixtures) GoRunConsume(t *testing.T, ctx context.Context) {
	t.Helper()
	go func() {
		err := consumer.Run(ctx, &kafka.ConfigKafka{
			Brokers:         s.brokers,
			Topic:           eventsTopic,
			ConsumerGroupID: uuid.New().String(),
		}, s.logger)
		assert.NoError(t, err)
	}()
}
