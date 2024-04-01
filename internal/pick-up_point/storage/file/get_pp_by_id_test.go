package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pick-up_point/storage"
	"homework/tests/fixtures"
	"homework/tests/states"
)

func Test_GetPickUpPointByID(t *testing.T) {
	t.Parallel()
	s := setUp()
	ctx := context.Background()

	t.Run("not exists", func(t *testing.T) {
		t.Parallel()

		pp, err := s.st.GetPickUpPointByID(ctx, states.PPID+1)

		assert.Nil(t, pp)
		assert.Equal(t, storage.ErrPickUpPointNotFound, err)
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		pp, err := s.st.GetPickUpPointByID(ctx, states.PPID)

		assert.NoError(t, err)
		assert.Equal(t, fixtures.PickUpPoint().Valid().P(), pp)
	})
}
