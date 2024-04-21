//go:build integration
// +build integration

package pick_up_points

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework/tests/test_json"
)

func TestGetPickUpPoints(t *testing.T) {

	t.Run("ok", func(t *testing.T) {
		del := setUp(t, tableName)
		request := httptest.NewRequest(http.MethodGet, "/pick-up-points", nil)
		respWriter := httptest.NewRecorder()

		del.GetPickUpPoints(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, test_json.ValidPPSliceResponse, string(body))
	})
}
