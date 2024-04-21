package storage

import (
	"context"
)

func (s *OrderStoragePG) DeleteOrdersByPPID(ctx context.Context, ppID uint64) error {
	_, err := s.db.Exec(ctx, `DELETE FROM orders WHERE pick_up_point_id = $1`, ppID)
	return err
}
