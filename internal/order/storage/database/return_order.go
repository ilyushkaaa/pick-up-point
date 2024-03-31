package storage

import (
	"context"

	"homework/internal/order/storage"
)

func (s *OrderStoragePG) ReturnOrder(ctx context.Context, clientID, orderID uint64) error {
	result, err := s.db.Cluster.Exec(ctx, "UPDATE orders SET is_returned = $1 WHERE client_id = $2 AND id = $3",
		true, clientID, orderID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrClientOrderNotFound
	}
	return nil
}
