package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage"
	"homework/internal/pick-up_point/storage/database/dto"
)

func (s *PPStorageDB) GetPickUpPointByName(ctx context.Context, name string) (*model.PickUpPoint, error) {
	var ppDB dto.PickUpPointDB
	err := s.db.Get(ctx, &ppDB,
		`SELECT id, name, phone_number, region, city, street, house_num 
                FROM pick_up_points WHERE name = $1`, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrPickUpPointNotFound
		}
		return nil, err
	}
	pp := dto.ConvertToPickUpPoint(ppDB)
	return &pp, nil
}
