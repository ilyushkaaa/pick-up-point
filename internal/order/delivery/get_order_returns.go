package delivery

import (
	"fmt"
	"strconv"
)

func (od *OrderDelivery) GetOrderReturnsDelivery(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("bad number of params")
	}

	pageNum := 1
	if len(args) == 1 {
		var err error
		pageNum, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("page number must be integer: %w", err)
		}
	}

	orders, err := od.service.GetOrderReturnsService(pageNum)
	if err != nil {
		return fmt.Errorf("error in getting order returns: %w", err)
	}

	for _, order := range orders {
		fmt.Printf("%+v\n", order)
	}
	return nil
}
