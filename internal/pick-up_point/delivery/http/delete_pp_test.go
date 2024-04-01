package delivery

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/storage"
	"homework/tests/test_json"
)

func Test_DeletePickUpPoint(t *testing.T) {
	t.Parallel()

	t.Run("bad id passed", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		request := httptest.NewRequest(http.MethodDelete, "/pick-up-point/bad_id", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "bad_id"})
		respWriter := httptest.NewRecorder()

		s.del.DeletePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, `{"result":"pick-up point ID must be positive integer"}`, string(body))
	})

	t.Run("pick-up point with such id was not found", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		request := httptest.NewRequest(http.MethodDelete, "/pick-up-point/5000", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "5000"})
		s.mockService.EXPECT().DeletePickUpPoint(request.Context(), uint64(5000)).Return(storage.ErrPickUpPointNotFound)
		respWriter := httptest.NewRecorder()

		s.del.DeletePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.Equal(t, `{"result":"no pick-up points with such id"}`, string(body))
	})

	t.Run("internal error", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		request := httptest.NewRequest(http.MethodDelete, "/pick-up-point/5000", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "5000"})
		s.mockService.EXPECT().DeletePickUpPoint(request.Context(), uint64(5000)).Return(fmt.Errorf("internal error"))
		respWriter := httptest.NewRecorder()

		s.del.DeletePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.Equal(t, test_json.InternalError, string(body))
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		request := httptest.NewRequest(http.MethodDelete, "/pick-up-point/5000", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "5000"})
		s.mockService.EXPECT().DeletePickUpPoint(request.Context(), uint64(5000)).Return(nil)
		respWriter := httptest.NewRecorder()

		s.del.DeletePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, test_json.SuccessResult, string(body))
	})

}
