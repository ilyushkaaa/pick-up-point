package delivery

import (
	"fmt"
	"strconv"
)

func (od *OrderDelivery) DeleteOrderDelivery(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("bad number of params")
	}
	orderID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("order ID must be integer: %w", err)
	}
	err = od.service.DeleteOrderService(orderID)
	if err != nil {
		return fmt.Errorf("error in returning order to courier: %w", err)
	}
	fmt.Println("order was returned to courier")
	return nil
}
