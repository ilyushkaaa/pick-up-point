package delivery

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/command_pp/response"
	"homework/tests/fixtures"
	"homework/tests/json_body"
)

func Test_UpdatePickUpPoint(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("error bad number of params", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		resp := s.del.UpdatePickUpPoint(ctx, []string{})

		assert.Equal(t, response.Response{
			Err: fmt.Errorf("update pick-up point method must have 1 param"),
		}, resp)
	})

	t.Run("validation error", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		resp := s.del.UpdatePickUpPoint(ctx, []string{json_body.InValidPPRequest})

		assert.Error(t, resp.Err)
	})

	t.Run("error in updating", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockService.EXPECT().UpdatePickUpPoint(ctx, fixtures.PickUpPoint().Valid().V()).Return(fmt.Errorf("internal error"))
		resp := s.del.UpdatePickUpPoint(ctx, []string{json_body.ValidPPUpdateRequest})

		assert.Error(t, resp.Err)

	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockService.EXPECT().UpdatePickUpPoint(ctx, fixtures.PickUpPoint().Valid().V()).Return(nil)
		resp := s.del.UpdatePickUpPoint(ctx, []string{json_body.ValidPPUpdateRequest})

		assert.NoError(t, resp.Err)
	})

}
