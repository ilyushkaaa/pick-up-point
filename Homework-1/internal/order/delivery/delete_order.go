package delivery

import (
	"fmt"
	"strconv"
)

func (od *OrderDelivery) DeleteOrderDelivery(args []string) {
	if len(args) != 1 {
		fmt.Println("bad number of params")
		return
	}
	orderID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("order ID must be integer")
		return
	}
	err = od.service.DeleteOrderService(orderID)
	if err != nil {
		fmt.Printf("error in returning order to courier: %s\n", err)
		return
	}
	fmt.Println("order was returned to courier")
}
