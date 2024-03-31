package delivery

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"homework/internal/order/model"
)

func (od *OrderDelivery) AddOrderDelivery(ctx context.Context, args []string) error {
	argsLen := len(args)
	if argsLen != 5 && argsLen != 6 {
		return fmt.Errorf("bad number of params")
	}
	orderID, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("order ID must be positive integer: %w", err)
	}
	clientID, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("client ID must be positive integer: %w", err)
	}
	expireDate, err := time.Parse("2006-01-02", args[2])
	if err != nil {
		return fmt.Errorf("max days for order storage must be date: %w", err)
	}
	weight, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		return fmt.Errorf("weight must be float: %w", err)
	}
	if weight <= 0 {
		return fmt.Errorf("weight must be positive: %f is not", weight)
	}
	price, err := strconv.ParseFloat(args[4], 64)
	if err != nil {
		return fmt.Errorf("price must be float: %w", err)
	}
	if price <= 0 {
		return fmt.Errorf("price must be positive: %f is not", price)
	}
	var packageType string
	if argsLen == 6 {
		packageType = args[5]
	}

	newOrder := model.NewOrder(orderID, clientID, weight, price, expireDate, packageType)

	err = od.service.AddOrder(ctx, newOrder)
	if err != nil {
		return fmt.Errorf("error in adding order to pick-up point: %w", err)
	}
	fmt.Println("order is added")
	return nil
}
