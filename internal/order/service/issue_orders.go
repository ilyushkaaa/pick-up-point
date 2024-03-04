package service

func (op *OrderServicePP) IssueOrderService(orderIDs map[int]struct{}) error {
	orders, err := op.storage.GetOrders()
	if err != nil {
		return err
	}
	ordersCount := 0
	clientID := 0
	clientIDWasSet := false
	for _, order := range orders {
		if _, exists := orderIDs[order.ID]; !exists {
			continue
		}
		if clientIDWasSet && clientID != order.ClientID {
			return ErrOrdersOfDifferentClients
		}
		if order.IsIssued {
			return ErrOrderAlreadyIssued
		}
		clientIDWasSet = true
		clientID = order.ClientID
		ordersCount++
	}
	if len(orderIDs) != ordersCount {
		return ErrOrderNotFound
	}
	return op.storage.IssueOrdersStorage(orderIDs)

}
