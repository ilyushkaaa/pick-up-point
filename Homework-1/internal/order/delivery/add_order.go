package delivery

import (
	"fmt"
	"strconv"
	"time"
)

func (od *OrderDelivery) AddOrderDelivery(args []string) {
	if len(args) != 3 {
		fmt.Println("bad number of params")
		return
	}
	orderID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("order ID must be integer")
		return
	}
	clientID, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("client ID must be integer")
		return
	}
	expireDate, err := time.Parse("2006-01-02", args[2])
	if err != nil {
		fmt.Println("max days for order storage must be integer")
		return
	}
	err = od.service.AddOrderService(orderID, clientID, expireDate)
	if err != nil {
		fmt.Printf("error in adding order to pick-up point: %s\n", err)
		return
	}
	fmt.Println("order is added")
}
