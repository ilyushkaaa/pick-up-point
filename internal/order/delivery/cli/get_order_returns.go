package delivery

import (
	"context"
	"fmt"
	"strconv"
)

func (od *OrderDelivery) GetOrderReturnsDelivery(ctx context.Context, args []string) error {
	if len(args) > 2 {
		return fmt.Errorf("bad number of params")
	}

	ordersPerPage, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("orders per page number must be positive integer: %w", err)
	}

	pageNum := uint64(1)
	if len(args) == 2 {
		pageNum, err = strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			return fmt.Errorf("page number must be positive integer: %w", err)
		}
	}

	orders, err := od.service.GetOrderReturns(ctx, ordersPerPage, pageNum)
	if err != nil {
		return fmt.Errorf("error in getting order returns: %w", err)
	}

	for _, order := range orders {
		fmt.Printf("%+v\n", order)
	}
	return nil
}
