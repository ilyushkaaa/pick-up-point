package service

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (ps *PPService) UpdatePickUpPoint(ctx context.Context, point model.PickUpPoint) error {
	wasUpdated, err := ps.storage.UpdatePickUpPoint(ctx, point)
	if err != nil {
		return err
	}
	if !wasUpdated {
		return ErrPickUpPointNotFound
	}
	return nil
}
