package pick_up_points

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/model"
)

func getPPFromResponse(t *testing.T, body []byte) model.PickUpPoint {
	var pp model.PickUpPoint
	err := json.Unmarshal(body, &pp)

	assert.NoError(t, err)
	return pp
}

func getPPSliceFromResponse(t *testing.T, body []byte) []model.PickUpPoint {
	var pp []model.PickUpPoint
	err := json.Unmarshal(body, &pp)

	assert.NoError(t, err)
	return pp
}
