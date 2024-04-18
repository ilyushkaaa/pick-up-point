package service

import (
	"context"
	"fmt"

	"homework/internal/pick-up_point/model"
)

func (ps *PPService) GetPickUpPointByID(ctx context.Context, id uint64) (*model.PickUpPoint, error) {
	gotFromCache := false
	var pickUpPoint *model.PickUpPoint
	data, err := ps.cache.GetFromCache(ctx, fmt.Sprintf("pp_%d", id))
	if err == nil {
		pickUpPoint, gotFromCache = data.(*model.PickUpPoint)
	}
	if !gotFromCache {
		pickUpPoint, err = ps.ppStorage.GetPickUpPointByID(ctx, id)
		if err != nil {
			return nil, err
		}
		ps.cache.GoAddToCache(context.Background(), fmt.Sprintf("pp_%d", id), pickUpPoint)
	}
	return pickUpPoint, nil
}
