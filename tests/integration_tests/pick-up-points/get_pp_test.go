package pick_up_points

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/model"
	"homework/tests/fixtures"
	"homework/tests/states"
)

func TestGetPickUpPoints(t *testing.T) {
	del, db := initTest(t)

	t.Run("ok", func(t *testing.T) {
		setUp(t, db, tableName)
		fillDataBase(t, db)
		ppExpected := []model.PickUpPoint{fixtures.PickUpPoint().Valid().V(),
			fixtures.PickUpPoint().Valid().ID(states.PPID2).Name(states.PPName2).V()}
		request := httptest.NewRequest(http.MethodGet, "/pick-up-points", nil)
		respWriter := httptest.NewRecorder()

		del.GetPickUpPoints(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		pp := getPPSliceFromResponse(t, body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, ppExpected, pp)
	})
}
