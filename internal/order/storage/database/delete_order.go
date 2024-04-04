package storage

import (
	"context"

	"homework/internal/order/storage"
)

func (s *OrderStoragePG) DeleteOrder(ctx context.Context, orderID uint64) error {
	result, err := s.db.Exec(ctx, `DELETE FROM orders WHERE id = $1`, orderID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrOrderNotFound
	}
	return nil
}
