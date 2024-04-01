package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/tests/fixtures"
)

func Test_AddPickUpPoint(t *testing.T) {
	t.Parallel()
	s := setUp()
	ctx := context.Background()

	t.Run("basic test", func(t *testing.T) {
		pp, err := s.st.AddPickUpPoint(ctx, fixtures.PickUpPoint().ValidWithoutID().V())

		assert.NoError(t, err)
		assert.Equal(t, fixtures.PickUpPoint().Valid().P(), pp)
	})

}
