package storage

import "context"

func (fs *FileOrderStorage) DeleteOrder(ctx context.Context, orderID uint64) error {
	orders, err := fs.GetOrders(ctx)
	if err != nil {
		return err
	}
	for i := 0; i < len(orders); i++ {
		if orders[i].ID == orderID {
			orders = append(orders[:i], orders[i+1:]...)
			break
		}
	}
	return fs.writeOrders(orders)
}
