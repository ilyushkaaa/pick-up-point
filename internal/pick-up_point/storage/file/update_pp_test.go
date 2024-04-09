package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/storage"
	"homework/tests/fixtures"
)

func Test_UpdatePickUpPoint(t *testing.T) {
	t.Parallel()
	s := setUp()
	ctx := context.Background()

	t.Run("not exists", func(t *testing.T) {
		t.Parallel()

		err := s.st.UpdatePickUpPoint(ctx, fixtures.PickUpPoint().Valid().Name("unknown").V())

		assert.ErrorIs(t, storage.ErrPickUpPointNotFound, err)
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		err := s.st.UpdatePickUpPoint(ctx, fixtures.PickUpPoint().Valid().HouseNum("222").V())

		assert.NoError(t, err)
	})
}
