package delivery

import (
	"fmt"
	"strconv"
)

func (od *OrderDelivery) ReturnOrderDelivery(args []string) {
	if len(args) != 2 {
		fmt.Println("bad number of params")
		return
	}
	clientID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("client ID must be integer")
		return
	}
	orderID, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("order ID must be integer")
		return
	}
	err = od.service.ReturnOrderService(clientID, orderID)
	if err != nil {
		fmt.Printf("error in returning order from client to pick-up point: %s\n", err)
		return
	}
	fmt.Println("client's order was returned to pick-up point")
}
