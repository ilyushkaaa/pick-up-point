package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"homework/internal/order/model"
	"homework/internal/order/storage/database/dto"
)

func (s *OrderStoragePG) GetUserOrders(ctx context.Context, clientID uint64) ([]model.Order, error) {
	var ordersDB []dto.OrderDB
	err := s.db.Select(ctx, &ordersDB,
		`SELECT id, client_id, weight, price, package_type, storage_expiration_date, order_issue_date, is_returned, pick_up_point_id 
                FROM orders WHERE client_id = $1`, clientID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	return dto.ConvertSliceToOrders(ordersDB), nil
}
