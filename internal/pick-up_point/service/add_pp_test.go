package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/storage"
	"homework/tests/fixtures"
	"homework/tests/states"
)

func Test_AddPickUpPoint(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("internal error in getting pp by name", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockStorage.EXPECT().GetPickUpPointByName(ctx, states.PPName1).Return(nil, fmt.Errorf("internal error"))

		pp, err := s.srv.AddPickUpPoint(ctx, fixtures.PickUpPoint().ValidWithoutID().V())

		assert.Nil(t, pp)
		assert.Equal(t, err, fmt.Errorf("internal error"))
	})

	t.Run("no errors, pick-up point with such name exists", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockStorage.EXPECT().GetPickUpPointByName(ctx, states.PPName1).Return(fixtures.PickUpPoint().Valid().P(), nil)

		pp, err := s.srv.AddPickUpPoint(ctx, fixtures.PickUpPoint().ValidWithoutID().V())

		assert.Nil(t, pp)
		assert.ErrorIs(t, err, ErrPickUpPointAlreadyExists)
	})

	t.Run("internal error in adding pick-up point", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockStorage.EXPECT().GetPickUpPointByName(ctx, states.PPName1).Return(nil, storage.ErrPickUpPointNotFound)
		s.mockStorage.EXPECT().AddPickUpPoint(ctx, fixtures.PickUpPoint().ValidWithoutID().V()).Return(nil, fmt.Errorf("internal error"))

		pp, err := s.srv.AddPickUpPoint(ctx, fixtures.PickUpPoint().ValidWithoutID().V())

		assert.Nil(t, pp)
		assert.Equal(t, err, fmt.Errorf("internal error"))
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockStorage.EXPECT().GetPickUpPointByName(ctx, states.PPName1).Return(nil, storage.ErrPickUpPointNotFound)
		s.mockStorage.EXPECT().AddPickUpPoint(ctx, fixtures.PickUpPoint().ValidWithoutID().V()).Return(fixtures.PickUpPoint().Valid().P(), nil)

		pp, err := s.srv.AddPickUpPoint(ctx, fixtures.PickUpPoint().ValidWithoutID().V())

		assert.Equal(t, fixtures.PickUpPoint().Valid().P(), pp)
		assert.NoError(t, err)
	})

}
