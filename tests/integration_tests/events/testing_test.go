//go:build integration
// +build integration

package events

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	eventsProducer "homework/internal/events/service/producer"
	"homework/internal/middleware"
	"homework/pkg/kafka/producer"
)

type TestEventsFixtures struct {
	mw           *middleware.Middleware
	buf          *bytes.Buffer
	syncProducer *producer.SyncProducer
	brokers      []string
	logger       *zap.SugaredLogger
}

type fakeHandler struct{}

func (h *fakeHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func setUp(t *testing.T) TestEventsFixtures {
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

	ep := eventsProducer.New(syncProducer, "test_events")
	mw := middleware.New(zapL, ep)

	return TestEventsFixtures{
		mw:           mw,
		buf:          &buf,
		syncProducer: syncProducer,
		brokers:      brokers,
		logger:       zapL,
	}
}

func tearDown(t *testing.T, cancelFunc context.CancelFunc, syncProducer *producer.SyncProducer) {
	t.Helper()

	cancelFunc()
	err := syncProducer.Close()
	assert.NoError(t, err)
}
