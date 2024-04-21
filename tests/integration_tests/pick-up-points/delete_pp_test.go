//go:build integration
// +build integration

package pick_up_points

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework/tests/test_json"
)

func TestDeletePickUpPointBy(t *testing.T) {

	t.Run("error no pick-up points with such id", func(t *testing.T) {
		del := setUp(t, tableName)
		request := httptest.NewRequest(http.MethodDelete, "/pick-up-point/5020", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "5020"})
		respWriter := httptest.NewRecorder()

		del.DeletePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.JSONEq(t, `{"result":"no pick-up points with such id"}`, string(body))
	})

	t.Run("ok", func(t *testing.T) {
		del := setUp(t, tableName)
		request := httptest.NewRequest(http.MethodDelete, "/pick-up-point/5000", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "5000"})
		respWriter := httptest.NewRecorder()

		del.DeletePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, test_json.SuccessResult, string(body))
	})
}
