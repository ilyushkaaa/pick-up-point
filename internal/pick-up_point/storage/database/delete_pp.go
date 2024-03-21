package storage

import "context"

func (s *PPStorageDB) DeletePickUpPoint(ctx context.Context, id uint64) (bool, error) {
	result, err := s.cluster.Exec(ctx, `DELETE FROM pick_up_points WHERE id = $1`, id)
	if err != nil {
		return false, err
	}
	if result.RowsAffected() == 0 {
		return false, nil
	}
	return true, nil
}
