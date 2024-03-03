package delivery

import (
	"fmt"
	"strconv"
)

func (od *OrderDelivery) IssueOrderDelivery(args []string) {
	if len(args) == 0 {
		fmt.Println("bad number of params")
		return
	}
	orderIDs := make(map[int]struct{}, len(args))
	for _, arg := range args {
		orderID, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("order ID must be integer")
			return
		}
		if _, exists := orderIDs[orderID]; exists {
			fmt.Println("order IDs has duplicates")
			return
		}
		orderIDs[orderID] = struct{}{}
	}
	err := od.service.IssueOrderService(orderIDs)
	if err != nil {
		fmt.Printf("error in issuing order to client: %s\n", err)
		return
	}
	fmt.Println("orders were issued to client")
}
