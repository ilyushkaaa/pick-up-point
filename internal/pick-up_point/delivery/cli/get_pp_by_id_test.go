package delivery

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework/tests/fixtures"
)

func Test_GetPickUpPointByID(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("error bad number of params", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)

		resp := s.del.GetPickUpPointByID(ctx, []string{})

		require.Error(t, resp.Err)
		assert.Empty(t, resp.Body)
		assert.Equal(t, "bad input params number: it must be 1", resp.Err.Error())
	})

	t.Run("error id is not uint64", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)

		resp := s.del.GetPickUpPointByID(ctx, []string{"bad_id"})

		require.Error(t, resp.Err)
		assert.Empty(t, resp.Body)
		assert.Equal(t, `pick-up point ID "bad_id" must be positive integer`, resp.Err.Error())
	})

	t.Run("error id getting pick-up point by name", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockService.EXPECT().GetPickUpPointByID(ctx, uint64(1)).Return(nil, fmt.Errorf("internal error"))

		resp := s.del.GetPickUpPointByID(ctx, []string{"1"})

		require.Error(t, resp.Err)
		assert.Empty(t, resp.Body)
		assert.Equal(t, `error in getting pick-up point by id: internal error`, resp.Err.Error())
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockService.EXPECT().GetPickUpPointByID(ctx, uint64(1)).Return(fixtures.PickUpPoint().Valid().P(), nil)

		resp := s.del.GetPickUpPointByID(ctx, []string{"1"})

		assert.NoError(t, resp.Err)
		assert.Equal(t, "&{ID:5000 Name:PickUpPoint1 Address:{Region:Курская область City:Курск Street:Студенческая HouseNum:2A} PhoneNumber:88005553535}", resp.Body)
	})

}
