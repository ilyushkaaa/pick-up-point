package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/tests/fixtures"
)

func Test_UpdatePickUpPoint(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("error in updating pick-up points", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockStorage.EXPECT().UpdatePickUpPoint(ctx, fixtures.PickUpPoint().Valid().V()).Return(fmt.Errorf("internal error"))
		err := s.srv.UpdatePickUpPoint(ctx, fixtures.PickUpPoint().Valid().V())

		assert.Equal(t, err, fmt.Errorf("internal error"))
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockStorage.EXPECT().UpdatePickUpPoint(ctx, fixtures.PickUpPoint().Valid().V()).Return(nil)
		err := s.srv.UpdatePickUpPoint(ctx, fixtures.PickUpPoint().Valid().V())

		assert.NoError(t, err)
	})
}
