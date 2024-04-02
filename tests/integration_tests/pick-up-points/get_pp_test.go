package pick_up_points

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/tests/test_json"
)

func TestGetPickUpPoints(t *testing.T) {
	del, db := initTest(t)

	t.Run("ok", func(t *testing.T) {
		setUp(t, db, tableName)
		fillDataBase(t, db)
		request := httptest.NewRequest(http.MethodGet, "/pick-up-points", nil)
		respWriter := httptest.NewRecorder()

		del.GetPickUpPoints(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, test_json.ValidPPSliceResponse, string(body))
	})
}
