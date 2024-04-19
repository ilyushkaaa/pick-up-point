//go:build integration
// +build integration

package events

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	eventsProducer "homework/internal/events/service/producer"
	"homework/internal/middleware"
	"homework/pkg/infrastructure/kafka"
	"homework/pkg/infrastructure/kafka/consumer"
	producer2 "homework/pkg/infrastructure/kafka/producer"
)

const eventsTopic = "test_events"

type TestEventsFixtures struct {
	mw           *middleware.Middleware
	buf          *bytes.Buffer
	syncProducer *producer2.SyncProducer
	brokers      []string
	logger       *zap.SugaredLogger
	cancel       context.CancelFunc
	waitChan     chan struct{}
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
	syncProducer, err := producer2.New(brokers)
	require.NoError(t, err)

	t.Cleanup(func() {
		err = syncProducer.Close()
		require.NoError(t, err)
	})

	ep := eventsProducer.New(syncProducer, eventsTopic)
	mw := middleware.New(zapL, ep)

	ctx, cancel := context.WithCancel(context.Background())
	waitChan := make(chan struct{})

	go func() {
		ticker := time.Tick(time.Second)
		for range ticker {
			content := buf.String()
			fmt.Println(content)
			if strings.Contains(content, "New") {
				waitChan <- struct{}{}
				return
			}
		}
	}()

	testFixtures := TestEventsFixtures{
		mw:           mw,
		buf:          &buf,
		syncProducer: syncProducer,
		brokers:      brokers,
		logger:       zapL,
		cancel:       cancel,
		waitChan:     waitChan,
	}
	testFixtures.GoRunConsume(ctx, t, waitChan)

	return testFixtures
}

func (s TestEventsFixtures) GoRunConsume(ctx context.Context, t *testing.T, waitChan chan struct{}) {
	t.Helper()
	go func() {
		err := consumer.Run(ctx, &kafka.ConfigKafka{
			Brokers:         s.brokers,
			Topic:           eventsTopic,
			ConsumerGroupID: uuid.New().String(),
		}, s.logger, waitChan)
		assert.NoError(t, err)
	}()
}
