package service

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (ps *PPService) UpdatePickUpPoint(ctx context.Context, point model.PickUpPoint) error {
	return ps.ppStorage.UpdatePickUpPoint(ctx, point)
}
