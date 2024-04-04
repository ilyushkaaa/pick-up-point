package storage

import (
	"context"

	"homework/internal/pick-up_point/storage"
)

func (s *PPStorageDB) DeletePickUpPoint(ctx context.Context, id uint64) error {
	result, err := s.db.Exec(ctx, `DELETE FROM pick_up_points WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrPickUpPointNotFound
	}
	return nil
}
