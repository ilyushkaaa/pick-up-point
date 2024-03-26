package service

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (ps *PPService) GetPickUpPointByID(ctx context.Context, id uint64) (*model.PickUpPoint, error) {
	pickUpPoint, err := ps.storage.GetPickUpPointByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return pickUpPoint, nil
}
