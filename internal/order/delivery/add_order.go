package delivery

import (
	"fmt"
	"strconv"
	"time"
)

func (od *OrderDelivery) AddOrderDelivery(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("bad number of params")
	}
	orderID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("order ID must be integer: %w", err)
	}
	clientID, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("client ID must be integer: %w", err)
	}
	expireDate, err := time.Parse("2006-01-02", args[2])
	if err != nil {
		return fmt.Errorf("max days for order storage must be date: %w", err)
	}
	err = od.service.AddOrderService(orderID, clientID, expireDate)
	if err != nil {
		return fmt.Errorf("error in adding order to pick-up point: %s", err)
	}
	fmt.Println("order is added")
	return nil
}
