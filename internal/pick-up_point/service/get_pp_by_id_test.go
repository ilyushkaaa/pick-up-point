package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/tests/fixtures"
	"homework/tests/states"
)

func Test_GetPickUpPointByID(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("error in getting pp by id", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockStorage.EXPECT().GetPickUpPointByID(ctx, states.PPID).Return(nil, fmt.Errorf("internal error"))
		pp, err := s.srv.GetPickUpPointByID(ctx, states.PPID)

		assert.Nil(t, pp)
		assert.Equal(t, err, fmt.Errorf("internal error"))
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockStorage.EXPECT().GetPickUpPointByID(ctx, states.PPID).Return(fixtures.PickUpPoint().Valid().P(), nil)
		pp, err := s.srv.GetPickUpPointByID(ctx, states.PPID)

		assert.Equal(t, fixtures.PickUpPoint().Valid().P(), pp)
		assert.NoError(t, err)
	})
}
