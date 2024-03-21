package service

import "context"

func (ps *PPService) DeletePickUpPoint(ctx context.Context, id uint64) error {
	wasDeleted, err := ps.storage.DeletePickUpPoint(ctx, id)
	if err != nil {
		return err
	}
	if !wasDeleted {
		return ErrPickUpPointNotFound
	}
	return nil
}
