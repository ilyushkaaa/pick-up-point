package delivery

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework/internal/pick-up_point/service"
	"homework/tests/fixtures"
	"homework/tests/test_json"
)

func Test_AddPickUpPoint(t *testing.T) {
	t.Parallel()

	t.Run("error bad json", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(`{"`))
		respWriter := httptest.NewRecorder()

		s.del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.JSONEq(t, test_json.InvalidInput, string(body))

	})

	t.Run("error validation", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(test_json.InValidPPRequest))
		respWriter := httptest.NewRecorder()

		s.del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.JSONEq(t, `{"error":"Address.house_num: non zero value required"}`, string(body))
	})

	t.Run("error pick-up point already exists", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(test_json.ValidPPAddRequest))
		s.mockService.EXPECT().AddPickUpPoint(gomock.Any(), fixtures.PickUpPoint().ValidWithoutID().V()).Return(nil, service.ErrPickUpPointAlreadyExists)
		respWriter := httptest.NewRecorder()

		s.del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.JSONEq(t, `{"error":"pick-up point with such name already exists"}`, string(body))
	})

	t.Run("internal error", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(test_json.ValidPPAddRequest))
		s.mockService.EXPECT().AddPickUpPoint(gomock.Any(), fixtures.PickUpPoint().ValidWithoutID().V()).Return(nil, fmt.Errorf("internal error"))
		respWriter := httptest.NewRecorder()

		s.del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.JSONEq(t, test_json.InternalError, string(body))
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(test_json.ValidPPAddRequest))
		s.mockService.EXPECT().AddPickUpPoint(gomock.Any(), fixtures.PickUpPoint().ValidWithoutID().V()).Return(fixtures.PickUpPoint().Valid().P(), nil)
		respWriter := httptest.NewRecorder()

		s.del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, test_json.ValidPPResponse, string(body))
	})

}
