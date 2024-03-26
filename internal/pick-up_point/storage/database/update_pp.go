package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage"
)

func (s *PPStorageDB) UpdatePickUpPoint(ctx context.Context, point model.PickUpPoint) error {
	result, err := s.cluster.Exec(ctx, `UPDATE pick_up_points SET phone_number = $1, region = $2, 
                          city = $3, street = $4, house_num = $5, name = $6 WHERE id = $7`,
		point.PhoneNumber, point.Address.Region, point.Address.City, point.Address.Street, point.Address.HouseNum,
		point.Name, point.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return storage.ErrPickUpPointNameExists
		}
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrPickUpPointNotFound
	}
	return nil
}
