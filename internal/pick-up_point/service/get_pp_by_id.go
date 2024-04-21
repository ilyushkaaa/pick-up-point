package service

import (
	"context"
	"strconv"

	"homework/internal/pick-up_point/model"
)

func (ps *PPService) GetPickUpPointByID(ctx context.Context, id uint64) (*model.PickUpPoint, error) {
	gotFromCache := false
	var pickUpPoint *model.PickUpPoint
	data, err := ps.cache.GetFromCache(ctx, strconv.FormatUint(id, 10))
	if err == nil {
		pickUpPoint, gotFromCache = data.(*model.PickUpPoint)
	}
	if !gotFromCache {
		pickUpPoint, err = ps.ppStorage.GetPickUpPointByID(ctx, id)
		if err != nil {
			return nil, err
		}
		ps.cache.GoAddToCache(context.Background(), strconv.FormatUint(id, 10), pickUpPoint)
	}
	return pickUpPoint, nil
}
