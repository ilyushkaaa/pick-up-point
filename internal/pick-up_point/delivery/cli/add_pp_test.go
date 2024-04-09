package delivery

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework/tests/fixtures"
	"homework/tests/test_json"
)

func Test_AddPickUpPoint(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("error bad number of params", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)

		resp := s.del.AddPickUpPoint(ctx, []string{})

		require.Error(t, resp.Err)
		assert.Empty(t, resp.Body)
		assert.Equal(t, "add pick-up point method must have 1 param", resp.Err.Error())
	})

	t.Run("validation error", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)

		resp := s.del.AddPickUpPoint(ctx, []string{test_json.InValidPPRequest})

		require.Error(t, resp.Err)
		assert.Empty(t, resp.Body)
		assert.Equal(t, "Address.house_num: non zero value required", resp.Err.Error())
	})

	t.Run("error in adding", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockService.EXPECT().AddPickUpPoint(ctx, fixtures.PickUpPoint().ValidWithoutID().V()).Return(nil, fmt.Errorf("internal error"))

		resp := s.del.AddPickUpPoint(ctx, []string{test_json.ValidPPAddRequest})

		require.Error(t, resp.Err)
		assert.Empty(t, resp.Body)
		assert.Equal(t, "error in adding new pick-up point: internal error", resp.Err.Error())
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockService.EXPECT().AddPickUpPoint(ctx, fixtures.PickUpPoint().ValidWithoutID().V()).Return(fixtures.PickUpPoint().Valid().P(), nil)

		resp := s.del.AddPickUpPoint(ctx, []string{test_json.ValidPPAddRequest})

		assert.NoError(t, resp.Err)
		assert.Equal(t, "pick-up point was successfully added: &{ID:5000 Name:PickUpPoint1 Address:{Region:Курская область City:Курск Street:Студенческая HouseNum:2A} PhoneNumber:88005553535}", resp.Body)
	})

}
