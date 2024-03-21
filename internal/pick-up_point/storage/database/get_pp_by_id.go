package storage

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage/database/dto"
)

func (s *PPStorageDB) GetPickUpPointByID(ctx context.Context, id uint64) (*model.PickUpPoint, error) {
	var ppDB dto.PickUpPointDB
	err := pgxscan.Get(ctx, s.cluster, &ppDB,
		`SELECT id, name, phone_number, region, city, street, house_num 
                FROM pick_up_points WHERE id = $1`, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	pp := ppDB.ConvertToPickUpPoint()
	return &pp, nil
}
