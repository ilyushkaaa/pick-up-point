package service

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (ps *PPService) GetPickUpPointByID(ctx context.Context, ID uint64) (*model.PickUpPoint, error) {
	pickUpPoint, err := ps.storage.GetPickUpPointByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	if pickUpPoint == nil {
		return nil, ErrPickUpPointNotFound
	}
	return pickUpPoint, nil
}
