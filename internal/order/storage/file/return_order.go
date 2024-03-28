package storage

import "context"

func (fs *FileOrderStorage) ReturnOrder(ctx context.Context, clientID, orderID uint64) error {
	orders, err := fs.GetOrders(ctx)
	if err != nil {
		return err
	}
	for i, order := range orders {
		if clientID == order.ClientID && orderID == order.ID {
			orders[i].IsReturned = true
			break
		}
	}
	return fs.writeOrders(orders)
}
