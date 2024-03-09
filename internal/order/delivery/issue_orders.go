package delivery

import (
	"fmt"
	"strconv"
)

func (od *OrderDelivery) IssueOrderDelivery(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("bad number of params")
	}
	orderIDs := make(map[int]struct{}, len(args))
	for _, arg := range args {
		orderID, err := strconv.Atoi(arg)
		if err != nil {
			return fmt.Errorf("order ID must be integer: %w", err)
		}
		if _, exists := orderIDs[orderID]; exists {
			return fmt.Errorf("order IDs has duplicates")
		}
		orderIDs[orderID] = struct{}{}
	}
	err := od.service.IssueOrderService(orderIDs)
	if err != nil {
		return fmt.Errorf("error in issuing order to client: %w", err)
	}
	fmt.Println("orders were issued to client")
	return nil
}
