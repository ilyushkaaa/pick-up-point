package delivery

import (
	"context"
	"fmt"
	"strconv"

	"homework/internal/filters/model"
)

func (od *OrderDelivery) GetUserOrdersDelivery(ctx context.Context, args []string) error {
	if len(args) == 0 || len(args) > 3 {
		return fmt.Errorf("bad number of params")
	}
	clientID, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("client ID must be positive integer: %w", err)
	}

	filters, err := makeFilters(args[1:])
	if err != nil {
		return err
	}

	orders, err := od.service.GetUserOrders(ctx, clientID, filters)
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

	for _, par := range params {
		switch par {
		case "PP-only":
			if filters.PPOnly {
				return filters, fmt.Errorf("can not set PP-only twice")
			}
			filters.PPOnly = true
		default:
			if filters.Limit != 0 {
				return filters, fmt.Errorf("can not set limit twice")
			}
			filters.Limit, err = parseLimit(params[0])
			if err != nil {
				return filters, err
			}
		}
	}
	return filters, nil
}
