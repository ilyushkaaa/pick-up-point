package service

import (
	"context"
	"fmt"

	cache "homework/internal/cache/in_memory"
	"homework/internal/pick-up_point/model"
)

func (ps *PPService) GetPickUpPointByID(ctx context.Context, id uint64) (*model.PickUpPoint, error) {
	gotFromCache := false
	var pickUpPoint *model.PickUpPoint
	data, err := ps.cache.GetFromCache(ctx, fmt.Sprintf("%s_%d", cache.PrefixPPByID, id))
	if err == nil {
		pickUpPoint, gotFromCache = data.(*model.PickUpPoint)
	}
	if !gotFromCache {
		pickUpPoint, err = ps.ppStorage.GetPickUpPointByID(ctx, id)
		if err != nil {
			return nil, err
		}
		ps.cache.GoAddToCache(context.Background(), fmt.Sprintf("%s_%d", cache.PrefixPPByID, id), pickUpPoint)
	}
	return pickUpPoint, nil
}
