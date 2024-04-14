package events

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoggingEvents(t *testing.T) {
	s := setUpAndConsume(t)
	defer s.cancelFunc()

	t.Run("test get", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://127.0.0.1/pick-up-points", nil)
		recorder := httptest.NewRecorder()
		time.Sleep(time.Second * 7)

		s.mw.AccessLog(&fakeHandler{}).ServeHTTP(recorder, req)
		time.Sleep(time.Second)

		assert.Contains(t, s.buf.String(), "New request")
		assert.Contains(t, s.buf.String(), `"method": "GET"`)
		assert.Contains(t, s.buf.String(), `"url": "/pick-up-points"`)

	})

	t.Run("test post", func(t *testing.T) {
		req := httptest.NewRequest("POST", "http://127.0.0.1/pick-up-point", nil)
		recorder := httptest.NewRecorder()
		time.Sleep(time.Second)

		s.mw.AccessLog(&fakeHandler{}).ServeHTTP(recorder, req)
		time.Sleep(time.Second)

		assert.Contains(t, s.buf.String(), "New request")
		assert.Contains(t, s.buf.String(), `"method": "POST"`)
		assert.Contains(t, s.buf.String(), `"url": "/pick-up-point"`)

	})

}
