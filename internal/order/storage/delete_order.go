package storage

func (fs *FileOrderStorage) DeleteOrderStorage(orderID int) error {
	orders, err := fs.GetOrders()
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
