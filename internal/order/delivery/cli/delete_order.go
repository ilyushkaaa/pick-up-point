package delivery

import (
	"context"
	"fmt"
	"strconv"
)

func (od *OrderDelivery) DeleteOrderDelivery(ctx context.Context, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("bad number of params")
	}
	orderID, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("order ID must be positive integer: %w", err)
	}
	err = od.service.DeleteOrder(ctx, orderID)
	if err != nil {
		return fmt.Errorf("error in returning order to courier: %w", err)
	}
	fmt.Println("order was returned to courier")
	return nil
}
