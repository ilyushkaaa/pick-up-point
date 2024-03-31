package storage

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"homework/internal/order/model"
	"homework/internal/order/storage/database/dto"
)

func (s *OrderStoragePG) GetUserOrders(ctx context.Context, clientID uint64) ([]model.Order, error) {
	var ordersDB []dto.OrderDB
	err := pgxscan.Select(ctx, s.db.Cluster, &ordersDB,
		`SELECT id, client_id, weight, price, package_type, storage_expiration_date, order_issue_date, is_returned 
                FROM orders WHERE client_id = $1`, clientID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	return dto.ConvertSliceToOrders(ordersDB), nil
}
