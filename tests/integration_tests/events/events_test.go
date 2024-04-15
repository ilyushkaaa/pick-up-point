//go:build integration
// +build integration

package events

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework/pkg/kafka/consumer"
)

const testTimeout = time.Second * 7

func TestLoggingEvents(t *testing.T) {

	t.Run("test get", func(t *testing.T) {
		s := setUpAndConsume(t)
		defer s.cancel()
		req := httptest.NewRequest("GET", "http://127.0.0.1/pick-up-points", nil)
		recorder := httptest.NewRecorder()
		require.NoError(t, consumer.WaitForConsumerReady(s.waitChan))

		s.mw.AccessLog(&fakeHandler{}).ServeHTTP(recorder, req)

		select {
		case <-s.waitChan:
			assert.Contains(t, s.buf.String(), `"method": "GET"`)
			assert.Contains(t, s.buf.String(), `"url": "/pick-up-points"`)
		case <-time.After(testTimeout):
			t.Error("Timeout occurred")
		}

	})

	t.Run("test post", func(t *testing.T) {
		s := setUpAndConsume(t)
		defer s.cancel()
		req := httptest.NewRequest("POST", "http://127.0.0.1/pick-up-point", nil)
		recorder := httptest.NewRecorder()
		require.NoError(t, consumer.WaitForConsumerReady(s.waitChan))

		s.mw.AccessLog(&fakeHandler{}).ServeHTTP(recorder, req)

		select {
		case <-s.waitChan:
			assert.Contains(t, s.buf.String(), `"method": "POST"`)
			assert.Contains(t, s.buf.String(), `"url": "/pick-up-point"`)
		case <-time.After(testTimeout):
			t.Error("Timeout occurred")
		}
	})

}
