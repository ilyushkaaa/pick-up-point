//go:build integration
// +build integration

package events

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"homework/pkg/kafka/consumer"
)

func TestLoggingEvents(t *testing.T) {

	t.Run("ok", func(t *testing.T) {
		s := setUp(t)
		ctx, cancel := context.WithCancel(context.Background())
		defer tearDown(t, cancel, s.syncProducer)

		go func() {
			err := consumer.Run(s.brokers, s.logger, ctx, "test_events", "test_group")
			assert.NoError(t, err)
		}()

		req := httptest.NewRequest("GET", "http://127.0.0.1/pick-up-points", nil)
		recorder := httptest.NewRecorder()
		time.Sleep(time.Second * 7)
		s.mw.AccessLog(&fakeHandler{}).ServeHTTP(recorder, req)
		time.Sleep(time.Second)
		fmt.Println(s.buf.String())
		assert.Contains(t, s.buf.String(), "New request")
		assert.Contains(t, s.buf.String(), `"method": "GET"`)
		assert.Contains(t, s.buf.String(), `"url": "/pick-up-points"`)
	})

}
