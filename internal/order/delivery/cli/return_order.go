package delivery

import (
	"context"
	"fmt"
	"strconv"
)

func (od *OrderDelivery) ReturnOrderDelivery(ctx context.Context, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("bad number of params")
	}

	clientID, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("client ID must be positive integer: %w", err)
	}

	orderID, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("order ID must be positive integer: %w", err)
	}
	err = od.service.ReturnOrder(ctx, clientID, orderID)
	if err != nil {
		return fmt.Errorf("error in returning order from client to pick-up point: %w", err)
	}
	fmt.Println("client's order was returned to pick-up point")
	return nil
}
