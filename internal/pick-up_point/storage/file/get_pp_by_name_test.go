package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/storage"
	"homework/tests/fixtures"
	"homework/tests/states"
)

func Test_GetPickUpPointByName(t *testing.T) {
	t.Parallel()
	s := setUp()
	ctx := context.Background()

	t.Run("not exists", func(t *testing.T) {
		t.Parallel()

		pp, err := s.st.GetPickUpPointByName(ctx, "unknown")

		assert.Nil(t, pp)
		assert.Equal(t, storage.ErrPickUpPointNotFound, err)
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		pp, err := s.st.GetPickUpPointByName(ctx, states.PPName1)

		assert.NoError(t, err)
		assert.Equal(t, fixtures.PickUpPoint().Valid().P(), pp)
	})
}
