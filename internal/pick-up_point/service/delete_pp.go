package service

import "context"

func (ps *PPService) DeletePickUpPoint(ctx context.Context, ID uint64) error {
	wasDeleted, err := ps.storage.DeletePickUpPoint(ctx, ID)
	if err != nil {
		return err
	}
	if !wasDeleted {
		return ErrPickUpPointNotFound
	}
	return nil
}
