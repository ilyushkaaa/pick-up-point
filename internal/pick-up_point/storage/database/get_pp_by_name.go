package storage

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage/database/dto"
)

func (s *PPStorageDB) GetPickUpPointByName(ctx context.Context, name string) (*model.PickUpPoint, error) {
	var ppDB dto.PickUpPointDB
	err := pgxscan.Get(ctx, s.cluster, &ppDB,
		`SELECT id, name, phone_number, region, city, street, house_num 
                FROM pick_up_points WHERE name = $1`, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	pp := ppDB.ConvertToPickUpPoint()
	return &pp, nil
}
