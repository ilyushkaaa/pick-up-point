package service

import (
	"context"
	"strconv"

	"homework/internal/pick-up_point/model"
)

func (ps *PPService) UpdatePickUpPoint(ctx context.Context, point model.PickUpPoint) error {
	err := ps.ppStorage.UpdatePickUpPoint(ctx, point)
	if err != nil {
		ps.cache.GoDeleteFromCache(context.Background(), strconv.FormatUint(point.ID, 10))
	}
	return err
}
