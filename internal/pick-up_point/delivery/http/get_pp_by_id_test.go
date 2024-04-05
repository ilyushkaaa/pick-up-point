package delivery

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework/internal/pick-up_point/storage"
	"homework/tests/fixtures"
	"homework/tests/test_json"
)

func Test_GetPickUpPointByID(t *testing.T) {
	t.Parallel()

	t.Run("bad id passed", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)

		request := httptest.NewRequest(http.MethodGet, "/pick-up-point/bad_id", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "bad_id"})
		respWriter := httptest.NewRecorder()

		s.del.GetPickUpPointByID(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.JSONEq(t, `{"result":"pick-up point ID must be positive integer"}`, string(body))
	})

	t.Run("pick-up point with such id was not found", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		request := httptest.NewRequest(http.MethodGet, "/pick-up-point/5000", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "5000"})
		s.mockService.EXPECT().GetPickUpPointByID(gomock.Any(), uint64(5000)).Return(nil, storage.ErrPickUpPointNotFound)
		respWriter := httptest.NewRecorder()

		s.del.GetPickUpPointByID(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.JSONEq(t, `{"result":"no pick-up points with such id"}`, string(body))
	})

	t.Run("internal error", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		request := httptest.NewRequest(http.MethodGet, "/pick-up-point/5000", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "5000"})
		s.mockService.EXPECT().GetPickUpPointByID(gomock.Any(), uint64(5000)).Return(nil, fmt.Errorf("internal error"))
		respWriter := httptest.NewRecorder()

		s.del.GetPickUpPointByID(respWriter, request)
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
		request := httptest.NewRequest(http.MethodGet, "/pick-up-point/5000", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "5000"})
		s.mockService.EXPECT().GetPickUpPointByID(gomock.Any(), uint64(5000)).Return(fixtures.PickUpPoint().Valid().P(), nil)
		respWriter := httptest.NewRecorder()

		s.del.GetPickUpPointByID(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, test_json.ValidPPResponse, string(body))
	})
}
