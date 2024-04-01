package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/model"
	"homework/tests/fixtures"
)

func Test_GetPickUpPoints(t *testing.T) {
	t.Parallel()
	s := setUp()
	ctx := context.Background()

	t.Run("basic test", func(t *testing.T) {
		pp, err := s.st.GetPickUpPoints(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []model.PickUpPoint{fixtures.PickUpPoint().Valid().V()}, pp)
	})

}
