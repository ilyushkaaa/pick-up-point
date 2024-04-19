package service

import (
	"context"
	"fmt"

	cache "homework/internal/cache/in_memory"
	"homework/internal/pick-up_point/model"
)

func (ps *PPService) UpdatePickUpPoint(ctx context.Context, point model.PickUpPoint) error {
	err := ps.ppStorage.UpdatePickUpPoint(ctx, point)
	if err != nil {
		ps.cache.GoDeleteFromCache(context.Background(), fmt.Sprintf("%s_%d", cache.PrefixPPByID, point.ID))
	}
	return err
}
