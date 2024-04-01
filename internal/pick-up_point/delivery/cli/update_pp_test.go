package delivery

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/tests/fixtures"
	"homework/tests/test_json"
)

func Test_UpdatePickUpPoint(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("error bad number of params", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		resp := s.del.UpdatePickUpPoint(ctx, []string{})

		assert.Error(t, resp.Err)
		assert.Empty(t, resp.Body)
		assert.Equal(t, "update pick-up point method must have 1 param", resp.Err.Error())
	})

	t.Run("validation error", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		resp := s.del.UpdatePickUpPoint(ctx, []string{test_json.InValidPPRequest})

		assert.Error(t, resp.Err)
		assert.Empty(t, resp.Body)
		assert.Equal(t, "Address.house_num: non zero value required;ID: non zero value required", resp.Err.Error())
	})

	t.Run("error in updating", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		s.mockService.EXPECT().UpdatePickUpPoint(ctx, fixtures.PickUpPoint().Valid().V()).Return(fmt.Errorf("internal error"))

		resp := s.del.UpdatePickUpPoint(ctx, []string{test_json.ValidPPUpdateRequest})

		assert.Error(t, resp.Err)
		assert.Empty(t, resp.Body)
		assert.Equal(t, "error in updating pick-up point: internal error", resp.Err.Error())
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()
		s.mockService.EXPECT().UpdatePickUpPoint(ctx, fixtures.PickUpPoint().Valid().V()).Return(nil)

		resp := s.del.UpdatePickUpPoint(ctx, []string{test_json.ValidPPUpdateRequest})

		assert.NoError(t, resp.Err)
		assert.Equal(t, "pick-up point was successfully updated", resp.Body)
	})

}
