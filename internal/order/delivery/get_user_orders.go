package delivery

import (
	"fmt"
	"strconv"

	"homework/internal/filters/model"
)

func (od *OrderDelivery) GetUserOrdersDelivery(args []string) error {
	if len(args) == 0 || len(args) > 3 {
		return fmt.Errorf("bad number of params")
	}
	clientID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("client ID must be integer: %w", err)
	}

	filters, err := makeFilters(args[1:])
	if err != nil {
		return err
	}

	orders, err := od.service.GetUserOrdersService(clientID, filters)
	if err != nil {
		return fmt.Errorf("error in getting user orders: %w", err)
	}
	if len(orders) == 0 {
		return fmt.Errorf("no orders found")
	}
	fmt.Printf("orders for client %d:\n", clientID)
	for i := 0; i < len(orders); i++ {
		fmt.Printf("%d. %+v\n", i+1, orders[i])
	}
	return nil
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

func makeFilters(params []string) (model.Filters, error) {
	filters := model.Filters{
		PPOnly: false,
		Limit:  0,
	}

	var err error

	switch len(params) {
	case 1:
		if params[0] == "PP-only" {
			filters.PPOnly = true
		} else {
			filters.Limit, err = parseLimit(params[0])
			if err != nil {
				return filters, err
			}
		}
	case 2:
		filters.Limit, err = parseLimit(params[0])
		if err != nil {
			return filters, err
		}
		if params[1] != "PP-only" {
			return filters, fmt.Errorf("unknown param")
		}
		filters.PPOnly = true
	}
	return filters, nil
}
