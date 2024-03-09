package storage

func (fs *FileOrderStorage) ReturnOrderStorage(clientID, orderID int) error {
	orders, err := fs.GetOrders()
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
