package service

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (ps *PPService) GetPickUpPoints(ctx context.Context) ([]model.PickUpPoint, error) {
	pickUpPoints, err := ps.storage.GetPickUpPoints(ctx)
	if err != nil {
		return nil, err
	}
	return pickUpPoints, nil
}
