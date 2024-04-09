package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/model"
	"homework/tests/fixtures"
)

func Test_GetPickUpPoints(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("error in getting pick-up points", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockStorage.EXPECT().GetPickUpPoints(ctx).Return(nil, fmt.Errorf("internal error"))

		pp, err := s.srv.GetPickUpPoints(ctx)

		assert.Nil(t, pp)
		assert.Equal(t, err, fmt.Errorf("internal error"))
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockStorage.EXPECT().GetPickUpPoints(ctx).Return([]model.PickUpPoint{fixtures.PickUpPoint().Valid().V()}, nil)

		pp, err := s.srv.GetPickUpPoints(ctx)

		assert.Equal(t, []model.PickUpPoint{fixtures.PickUpPoint().Valid().V()}, pp)
		assert.NoError(t, err)
	})
}
