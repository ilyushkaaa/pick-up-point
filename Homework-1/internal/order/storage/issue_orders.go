package storage

import (
	"time"
)

func (fs *FileOrderStorage) IssueOrdersStorage(orderIDs map[int]struct{}) error {
	orders, err := fs.GetOrders()
	if err != nil {
		return err
	}
	for i, ord := range orders {
		if _, exists := orderIDs[ord.ID]; exists {
			orders[i].IsIssued = true
			orders[i].OrderIssueDate = time.Now()
		}
	}
	return fs.writeOrders(orders)
}
