package delivery

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/model"
	"homework/tests/fixtures"
	"homework/tests/test_json"
)

func Test_GetPickUpPoints(t *testing.T) {
	t.Parallel()

	t.Run("internal error", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		request := httptest.NewRequest(http.MethodGet, "/pick-up-points", nil)
		s.mockService.EXPECT().GetPickUpPoints(request.Context()).Return(nil, fmt.Errorf("internal error"))
		respWriter := httptest.NewRecorder()

		s.del.GetPickUpPoints(respWriter, request)
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
		request := httptest.NewRequest(http.MethodGet, "/pick-up-points", nil)
		s.mockService.EXPECT().GetPickUpPoints(request.Context()).Return([]model.PickUpPoint{fixtures.PickUpPoint().Valid().V()}, nil)
		respWriter := httptest.NewRecorder()

		s.del.GetPickUpPoints(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, `[{"ID":5000,"Name":"PickUpPoint1","Address":{"Region":"Курская область","City":"Курск","Street":"Студенческая","HouseNum":"2A"},"PhoneNumber":"88005553535"}]`, string(body))
	})
}
