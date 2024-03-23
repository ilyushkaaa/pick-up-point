package service

import (
	"context"
	"errors"

	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage"
)

func (ps *PPService) UpdatePickUpPoint(ctx context.Context, point model.PickUpPoint) error {
	pp, err := ps.storage.GetPickUpPointByName(ctx, point.Name)
	if err != nil {
		if !errors.Is(err, storage.ErrPickUpPointNotFound) {
			return err
		}
	}
	if err == nil && pp != nil {
		return ErrPickUpPointAlreadyExists
	}
	return ps.storage.UpdatePickUpPoint(ctx, point)
}
