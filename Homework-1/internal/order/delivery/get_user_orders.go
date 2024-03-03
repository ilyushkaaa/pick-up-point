package delivery

import (
	"fmt"
	"strconv"
)

func (od *OrderDelivery) GetUserOrdersDelivery(args []string) {
	if len(args) == 0 || len(args) > 3 {
		fmt.Println("bad number of params")
		return
	}
	clientID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("client ID must be integer")
		return
	}
	ppOnly := false
	limit := 0

	if len(args) == 2 {
		if args[1] == "PP-only" {
			ppOnly = true
		} else {
			limit, err = parseLimit(args[1])
			if err != nil {
				return
			}
		}
	}
	if len(args) == 3 {
		limit, err = parseLimit(args[1])
		if err != nil {
			return
		}
		if args[2] != "PP-only" {
			fmt.Println("unknown param")
			return
		}
		ppOnly = true
	}

	orders, err := od.service.GetUserOrdersService(clientID, limit, ppOnly)
	if err != nil {
		fmt.Printf("error in getting user orders: %s\n", err)
		return
	}
	if len(orders) == 0 {
		fmt.Println("no orders found")
		return
	}
	fmt.Printf("orders for client %d:\n", clientID)
	for i := 0; i < len(orders); i++ {
		fmt.Printf("%d. %+v\n", i+1, orders[i])
	}
}

func parseLimit(lim string) (int, error) {
	limit, err := strconv.Atoi(lim)
	if err != nil {
		fmt.Println("limit must be integer")
		return 0, err
	}
	if limit < 1 {
		fmt.Println("limit must be 1 or more")
		return 0, fmt.Errorf("limit must be 1 or more")
	}
	return limit, nil
}
