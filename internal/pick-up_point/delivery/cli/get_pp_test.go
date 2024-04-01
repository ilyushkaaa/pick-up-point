package delivery

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/command_pp/response"
	"homework/internal/pick-up_point/model"
	"homework/tests/fixtures"
)

func Test_GetPickUpPoints(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("error bad number of params", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		resp := s.del.GetPickUpPoints(ctx, []string{"dummy"})

		assert.Equal(t, response.Response{
			Err: fmt.Errorf("this request must not contain any params"),
		}, resp)
	})

	t.Run("error in getting pick-up points", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockService.EXPECT().GetPickUpPoints(ctx).Return(nil, fmt.Errorf("internal error"))
		resp := s.del.GetPickUpPoints(ctx, []string{})

		assert.Error(t, resp.Err)

	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockService.EXPECT().GetPickUpPoints(ctx).Return([]model.PickUpPoint{fixtures.PickUpPoint().Valid().V()}, nil)
		resp := s.del.GetPickUpPoints(ctx, []string{})

		assert.NoError(t, resp.Err)
	})
}
