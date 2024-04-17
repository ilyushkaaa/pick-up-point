package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"homework/internal/order/model"
	"homework/internal/order/storage"
	"homework/internal/order/storage/database/dto"
)

func (s *OrderStoragePG) GetOrderByID(ctx context.Context, id uint64) (*model.Order, error) {
	var orderDB dto.OrderDB
	err := s.db.Get(ctx, &orderDB,
		`SELECT id, client_id, weight, price, package_type, storage_expiration_date, order_issue_date, is_returned, pick_up_point_id 
                FROM orders WHERE id = $1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrOrderNotFound
		}
		return nil, err
	}
	order := dto.ConvertToOrder(orderDB)
	return &order, nil
}
