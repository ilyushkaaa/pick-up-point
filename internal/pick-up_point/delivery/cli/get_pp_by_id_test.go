package delivery

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/command_pp/response"
	"homework/tests/fixtures"
)

func Test_GetPickUpPointByID(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("error bad number of params", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		resp := s.del.GetPickUpPointByID(ctx, []string{})

		assert.Equal(t, response.Response{
			Err: fmt.Errorf("bad input params number: it must be 1"),
		}, resp)
	})

	t.Run("error id is not uint64", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		resp := s.del.GetPickUpPointByID(ctx, []string{"bad_id"})

		assert.Equal(t, response.Response{
			Err: fmt.Errorf(`pick-up point ID "bad_id" must be positive integer`),
		}, resp)
	})

	t.Run("error id getting pick-up point by name", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockService.EXPECT().GetPickUpPointByID(ctx, uint64(1)).Return(nil, fmt.Errorf("internal error"))
		resp := s.del.GetPickUpPointByID(ctx, []string{"1"})

		assert.Error(t, resp.Err)
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockService.EXPECT().GetPickUpPointByID(ctx, uint64(1)).Return(fixtures.PickUpPoint().Valid().P(), nil)
		resp := s.del.GetPickUpPointByID(ctx, []string{"1"})

		assert.NoError(t, resp.Err)
	})

}
