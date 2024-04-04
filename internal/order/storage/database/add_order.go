package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"homework/internal/order/model"
	"homework/internal/order/storage"
	"homework/internal/order/storage/database/dto"
)

func (s *OrderStoragePG) AddOrder(ctx context.Context, newOrder model.Order) error {
	orderDB := dto.NewOrderDB(newOrder)
	_, err := s.db.Exec(ctx,
		`INSERT INTO orders (id, client_id, weight, price, package_type, storage_expiration_date, order_issue_date, is_returned)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		orderDB.ID, orderDB.ClientID, orderDB.Weight, orderDB.Price, orderDB.PackageType, orderDB.StorageExpirationDate, orderDB.OrderIssueDate, false)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return storage.ErrOrderAlreadyExists
		}
		return err
	}
	return nil
}
