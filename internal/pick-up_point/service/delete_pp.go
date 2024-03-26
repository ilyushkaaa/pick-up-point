package service

import "context"

func (ps *PPService) DeletePickUpPoint(ctx context.Context, id uint64) error {
	return ps.storage.DeletePickUpPoint(ctx, id)
}
