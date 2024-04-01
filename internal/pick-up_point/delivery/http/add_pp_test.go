package delivery

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/service"
	"homework/tests/fixtures"
	"homework/tests/json_body"
)

func Test_AddPickUpPoint(t *testing.T) {
	t.Parallel()

	t.Run("error bad json", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(`{"`))
		respWriter := httptest.NewRecorder()
		s.del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, json_body.InvalidInput, string(body))

	})

	t.Run("error validation", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(json_body.InValidPPRequest))
		respWriter := httptest.NewRecorder()
		s.del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, `{"result":"Address.house_num: non zero value required"}`, string(body))
	})

	t.Run("error pick-up point already exists", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(json_body.ValidPPAddRequest))
		s.mockService.EXPECT().AddPickUpPoint(request.Context(), fixtures.PickUpPoint().ValidWithoutID().V()).Return(nil, service.ErrPickUpPointAlreadyExists)
		respWriter := httptest.NewRecorder()
		s.del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, `{"result":"pick-up point with such name already exists"}`, string(body))
	})

	t.Run("internal error", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(json_body.ValidPPAddRequest))
		s.mockService.EXPECT().AddPickUpPoint(request.Context(), fixtures.PickUpPoint().ValidWithoutID().V()).Return(nil, fmt.Errorf("internal error"))
		respWriter := httptest.NewRecorder()
		s.del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.Equal(t, json_body.InternalError, string(body))
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(json_body.ValidPPAddRequest))
		s.mockService.EXPECT().AddPickUpPoint(request.Context(), fixtures.PickUpPoint().ValidWithoutID().V()).Return(fixtures.PickUpPoint().Valid().P(), nil)
		respWriter := httptest.NewRecorder()
		s.del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, json_body.ValidPPResponse, string(body))
	})

}
