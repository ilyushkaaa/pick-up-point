package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage"
	"homework/internal/pick-up_point/storage/database/dto"
)

func (s *PPStorageDB) UpdatePickUpPoint(ctx context.Context, point model.PickUpPoint) error {
	pointDB := dto.NewPickUpPointDB(point)
	result, err := s.db.Exec(ctx, `UPDATE pick_up_points SET phone_number = $1, region = $2, 
                          city = $3, street = $4, house_num = $5, name = $6 WHERE id = $7`,
		pointDB.PhoneNumber, pointDB.Region, pointDB.City, pointDB.Street, pointDB.HouseNum,
		pointDB.Name, pointDB.ID)
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
