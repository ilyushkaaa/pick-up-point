package delivery

import (
	"fmt"
	"strconv"
)

func (od *OrderDelivery) ReturnOrderDelivery(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("bad number of params")
	}
	clientID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("client ID must be integer: %w", err)
	}
	orderID, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("order ID must be integer: %w", err)
	}
	err = od.service.ReturnOrderService(clientID, orderID)
	if err != nil {
		return fmt.Errorf("error in returning order from client to pick-up point: %s", err)
	}
	fmt.Println("client's order was returned to pick-up point")
	return nil
}
