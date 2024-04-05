package delivery

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework/internal/pick-up_point/model"
	"homework/tests/fixtures"
	"homework/tests/test_json"
)

func Test_GetPickUpPoints(t *testing.T) {
	t.Parallel()

	t.Run("internal error", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		request := httptest.NewRequest(http.MethodGet, "/pick-up-points", nil)
		s.mockService.EXPECT().GetPickUpPoints(gomock.Any()).Return(nil, fmt.Errorf("internal error"))
		respWriter := httptest.NewRecorder()

		s.del.GetPickUpPoints(respWriter, request)
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
		request := httptest.NewRequest(http.MethodGet, "/pick-up-points", nil)
		s.mockService.EXPECT().GetPickUpPoints(gomock.Any()).Return([]model.PickUpPoint{fixtures.PickUpPoint().Valid().V()}, nil)
		respWriter := httptest.NewRecorder()

		s.del.GetPickUpPoints(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, `[{"ID":5000,"Name":"PickUpPoint1","Address":{"Region":"Курская область","City":"Курск","Street":"Студенческая","HouseNum":"2A"},"PhoneNumber":"88005553535"}]`, string(body))
	})
}
