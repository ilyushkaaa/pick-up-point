//go:build integration
// +build integration

package pick_up_points

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework/tests/test_json"
)

func TestUpdatePickUpPoint(t *testing.T) {
	del, db := initTest(t)

	t.Run("error pick-up point with such id was not found", func(t *testing.T) {
		setUp(t, db, tableName)
		fillDataBase(t, db)
		request := httptest.NewRequest(http.MethodPut, "/pick-up-point", strings.NewReader(test_json.ValidPPUpdateRequestNotExists))
		respWriter := httptest.NewRecorder()

		del.UpdatePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.JSONEq(t, `{"result":"no pick-up points with such id"}`, string(body))
	})

	t.Run("error pick-up point with such name already exists", func(t *testing.T) {
		setUp(t, db, tableName)
		fillDataBase(t, db)
		request := httptest.NewRequest(http.MethodPut, "/pick-up-point", strings.NewReader(test_json.ValidPPUpdateRequestNameAlreadyExists))
		respWriter := httptest.NewRecorder()

		del.UpdatePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.JSONEq(t, `{"result":"pick-up point with such name already exists"}`, string(body))
	})

	t.Run("ok", func(t *testing.T) {
		setUp(t, db, tableName)
		fillDataBase(t, db)
		request := httptest.NewRequest(http.MethodPut, "/pick-up-point", strings.NewReader(test_json.ValidPPUpdateRequest))
		respWriter := httptest.NewRecorder()

		del.UpdatePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, test_json.ValidPPResponse, string(body))
	})
}
