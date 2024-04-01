package delivery

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/storage"
	"homework/tests/fixtures"
	"homework/tests/test_json"
)

func Test_UpdatePickUpPoint(t *testing.T) {
	t.Parallel()

	t.Run("error validation", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(test_json.InValidPPRequest))
		respWriter := httptest.NewRecorder()

		s.del.UpdatePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, `{"result":"Address.house_num: non zero value required;ID: non zero value required"}`, string(body))
	})

	t.Run("error pick-up point with such id was not found", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(test_json.ValidPPUpdateRequest))
		s.mockService.EXPECT().UpdatePickUpPoint(request.Context(), fixtures.PickUpPoint().Valid().V()).Return(storage.ErrPickUpPointNotFound)
		respWriter := httptest.NewRecorder()

		s.del.UpdatePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.Equal(t, `{"result":"no pick-up points with such id"}`, string(body))
	})

	t.Run("error pick-up point with such name already exists", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(test_json.ValidPPUpdateRequest))
		s.mockService.EXPECT().UpdatePickUpPoint(request.Context(), fixtures.PickUpPoint().Valid().V()).Return(storage.ErrPickUpPointNameExists)
		respWriter := httptest.NewRecorder()

		s.del.UpdatePickUpPoint(respWriter, request)
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
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(test_json.ValidPPUpdateRequest))
		s.mockService.EXPECT().UpdatePickUpPoint(request.Context(), fixtures.PickUpPoint().Valid().V()).Return(fmt.Errorf("internal error"))
		respWriter := httptest.NewRecorder()

		s.del.UpdatePickUpPoint(respWriter, request)
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
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(test_json.ValidPPUpdateRequest))
		s.mockService.EXPECT().UpdatePickUpPoint(request.Context(), fixtures.PickUpPoint().Valid().V()).Return(nil)
		respWriter := httptest.NewRecorder()

		s.del.UpdatePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, test_json.ValidPPResponse, string(body))
	})

}
