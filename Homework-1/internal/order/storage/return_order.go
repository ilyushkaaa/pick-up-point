package storage

func (fs *FileOrderStorage) ReturnOrderStorage(clientID, orderID int) error {
	orders, err := fs.GetOrders()
	if err != nil {
		return err
	}
	for i, ord := range orders {
		if clientID == ord.ClientID && orderID == ord.ID {
			orders[i].IsReturned = true
			break
		}
	}
	return fs.writeOrders(orders)
}
