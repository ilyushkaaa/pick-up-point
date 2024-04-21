package service

import (
	"context"
	"errors"

	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage"
)

func (ps *PPService) AddPickUpPoint(ctx context.Context, point model.PickUpPoint) (*model.PickUpPoint, error) {
	pp, err := ps.ppStorage.GetPickUpPointByName(ctx, point.Name)
	if err != nil {
		if !errors.Is(err, storage.ErrPickUpPointNotFound) {
			return nil, err
		}
	}
	if err == nil && pp != nil {
		return nil, ErrPickUpPointAlreadyExists
	}
	addedPP, err := ps.ppStorage.AddPickUpPoint(ctx, point)
	if err != nil {
		return nil, err
	}
	return addedPP, nil
}
