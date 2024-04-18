package service

import (
	"context"
	"fmt"

	"homework/internal/pick-up_point/model"
)

func (ps *PPService) UpdatePickUpPoint(ctx context.Context, point model.PickUpPoint) error {
	err := ps.ppStorage.UpdatePickUpPoint(ctx, point)
	if err != nil {
		ps.cache.GoDeleteFromCache(context.Background(), fmt.Sprintf("pp_%d", point.ID))

	}
	return err
}
