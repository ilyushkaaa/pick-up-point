package storage

import (
	"context"
	"time"
)

func (fs *FileOrderStorage) IssueOrders(ctx context.Context, orderIDs map[uint64]struct{}) error {
	orders, err := fs.GetOrders(ctx)
	if err != nil {
		return err
	}
	for i, order := range orders {
		if _, exists := orderIDs[order.ID]; exists {
			orders[i].IsIssued = true
			orders[i].OrderIssueDate = time.Now()
		}
	}
	return fs.writeOrders(orders)
}
