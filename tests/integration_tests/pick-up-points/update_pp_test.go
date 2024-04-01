package pick_up_points

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/tests/fixtures"
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

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.Equal(t, `{"result":"no pick-up points with such id"}`, string(body))
	})

	t.Run("error pick-up point with such name already exists", func(t *testing.T) {
		setUp(t, db, tableName)
		fillDataBase(t, db)
		request := httptest.NewRequest(http.MethodPut, "/pick-up-point", strings.NewReader(test_json.ValidPPUpdateRequestNameAlreadyExists))
		respWriter := httptest.NewRecorder()

		del.UpdatePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, `{"result":"pick-up point with such name already exists"}`, string(body))
	})

	t.Run("ok", func(t *testing.T) {
		setUp(t, db, tableName)
		fillDataBase(t, db)
		ppExpected := fixtures.PickUpPoint().Valid().V()
		request := httptest.NewRequest(http.MethodPut, "/pick-up-point", strings.NewReader(test_json.ValidPPUpdateRequest))
		respWriter := httptest.NewRecorder()

		del.UpdatePickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		pp := getPPFromResponse(t, body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, ppExpected, pp)
	})
}
