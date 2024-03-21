package service

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (ps *PPService) AddPickUpPoint(ctx context.Context, point model.PickUpPoint) (*model.PickUpPoint, error) {
	pp, err := ps.storage.GetPickUpPointByName(ctx, point.Name)
	if err != nil {
		return nil, err
	}
	if pp != nil {
		return nil, ErrPickUpPointAlreadyExists
	}
	addedPP, err := ps.storage.AddPickUpPoint(ctx, point)
	if err != nil {
		return nil, err
	}
	return addedPP, nil
}
