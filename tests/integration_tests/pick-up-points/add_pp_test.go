package pick_up_points

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/tests/fixtures"
	"homework/tests/states"
	"homework/tests/test_json"
)

const tableName = "pick_up_points"

func TestAddPickUpPoint(t *testing.T) {
	del, db := initTest(t)

	t.Run("error pick-up point with such name already exists", func(t *testing.T) {
		setUp(t, db, tableName)
		fillDataBase(t, db)
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(test_json.ValidPPAddRequest))
		respWriter := httptest.NewRecorder()

		del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, `{"result":"pick-up point with such name already exists"}`, string(body))
	})

	t.Run("ok", func(t *testing.T) {
		setUp(t, db, tableName)
		fillDataBase(t, db)
		ppExpected := fixtures.PickUpPoint().Valid().Name(states.PPName3).ID(states.PPID3).V()
		request := httptest.NewRequest(http.MethodPost, "/pick-up-point", strings.NewReader(test_json.ValidPPAddRequestUnique))
		respWriter := httptest.NewRecorder()

		del.AddPickUpPoint(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		defer resp.Body.Close()
		pp := getPPFromResponse(t, body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, ppExpected, pp)
	})
}
