package pick_up_points

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"homework/tests/fixtures"
)

func TestGetPickUpPointByID(t *testing.T) {
	del, db := initTest(t)

	t.Run("error no pick-up points with such id", func(t *testing.T) {
		setUp(t, db, tableName)
		fillDataBase(t, db)
		request := httptest.NewRequest(http.MethodGet, "/pick-up-point/5020", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "5020"})
		respWriter := httptest.NewRecorder()

		del.GetPickUpPointByID(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.Equal(t, `{"result":"no pick-up points with such id"}`, string(body))
	})

	t.Run("ok", func(t *testing.T) {
		setUp(t, db, tableName)
		fillDataBase(t, db)
		ppExpected := fixtures.PickUpPoint().Valid().V()
		request := httptest.NewRequest(http.MethodGet, "/pick-up-point/5000", nil)
		request = mux.SetURLVars(request, map[string]string{"PP_ID": "5000"})
		respWriter := httptest.NewRecorder()

		del.GetPickUpPointByID(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		defer resp.Body.Close()
		pp := getPPFromResponse(t, body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, ppExpected, pp)
	})
}
