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

func Test_AddPickUpPoint(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("error bad number of params", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		resp := s.del.AddPickUpPoint(ctx, []string{})

		assert.Equal(t, response.Response{
			Err: fmt.Errorf("add pick-up point method must have 1 param"),
		}, resp)
	})

	t.Run("validation error", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		resp := s.del.AddPickUpPoint(ctx, []string{json_body.InValidPPRequest})

		assert.Error(t, resp.Err)
	})

	t.Run("error in adding", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockService.EXPECT().AddPickUpPoint(ctx, fixtures.PickUpPoint().ValidWithoutID().V()).Return(nil, fmt.Errorf("internal error"))
		resp := s.del.AddPickUpPoint(ctx, []string{json_body.ValidPPAddRequest})

		assert.Error(t, resp.Err)

	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockService.EXPECT().AddPickUpPoint(ctx, fixtures.PickUpPoint().ValidWithoutID().V()).Return(fixtures.PickUpPoint().Valid().P(), nil)
		resp := s.del.AddPickUpPoint(ctx, []string{json_body.ValidPPAddRequest})

		assert.NoError(t, resp.Err)
	})

}
