package delivery

import (
	"context"
	"fmt"
	"strconv"
)

func (od *OrderDelivery) IssueOrderDelivery(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("bad number of params")
	}
	orderIDs := make(map[uint64]struct{}, len(args))
	for _, arg := range args {
		orderID, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			return fmt.Errorf("order ID must be positive integer: %w", err)
		}
		if _, exists := orderIDs[orderID]; exists {
			return fmt.Errorf("order IDs has duplicates")
		}
		orderIDs[orderID] = struct{}{}
	}
	err := od.service.IssueOrders(ctx, orderIDs)
	if err != nil {
		return fmt.Errorf("error in issuing order to client: %w", err)
	}
	fmt.Println("orders were issued to client")
	return nil
}
