package storage

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"homework/internal/order/model"
	"homework/internal/order/storage/database/dto"
)

func (s *OrderStoragePG) GetOrders(ctx context.Context) ([]model.Order, error) {
	var ordersDB []dto.OrderDB
	err := pgxscan.Select(ctx, s.db.Cluster, &ordersDB,
		`SELECT id, client_id, weight, price, package_type, storage_expiration_date, order_issue_date, is_returned 
                FROM orders`)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	orders := make([]model.Order, len(ordersDB))
	for i := range ordersDB {
		orders[i] = ordersDB[i].ConvertToOrder()
	}
	return orders, nil
}
