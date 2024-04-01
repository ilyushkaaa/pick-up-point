package delivery

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

		assert.Error(t, resp.Err)
		assert.Empty(t, resp.Body)
		assert.Equal(t, "this request must not contain any params", resp.Err.Error())
	})

	t.Run("error in getting pick-up points", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		s.mockService.EXPECT().GetPickUpPoints(ctx).Return(nil, fmt.Errorf("internal error"))

		resp := s.del.GetPickUpPoints(ctx, []string{})

		assert.Error(t, resp.Err)
		assert.Empty(t, resp.Body)
		assert.Equal(t, "error in getting pick-up points: internal error", resp.Err.Error())
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		s.mockService.EXPECT().GetPickUpPoints(ctx).Return([]model.PickUpPoint{fixtures.PickUpPoint().Valid().V()}, nil)

		resp := s.del.GetPickUpPoints(ctx, []string{})

		assert.NoError(t, resp.Err)
		assert.Equal(t, "[{ID:5000 Name:PickUpPoint1 Address:{Region:Курская область City:Курск Street:Студенческая HouseNum:2A} PhoneNumber:88005553535}]", resp.Body)
	})
}
