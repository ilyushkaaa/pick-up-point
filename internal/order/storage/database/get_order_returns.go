package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"homework/internal/order/model"
	"homework/internal/order/storage/database/dto"
)

func (s *OrderStoragePG) GetOrderReturns(ctx context.Context) ([]model.Order, error) {
	var ordersDB []dto.OrderDB
	err := s.db.Select(ctx, &ordersDB,
		`SELECT id, client_id, weight, price, package_type, storage_expiration_date, order_issue_date, is_returned 
                FROM orders WHERE is_returned = true`)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	return dto.ConvertSliceToOrders(ordersDB), nil
}
