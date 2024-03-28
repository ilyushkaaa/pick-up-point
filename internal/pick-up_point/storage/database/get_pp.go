package storage

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage/database/dto"
)

func (s *PPStorageDB) GetPickUpPoints(ctx context.Context) ([]model.PickUpPoint, error) {
	var pickUpPointsDB []dto.PickUpPointDB
	err := pgxscan.Select(ctx, s.db.Cluster, &pickUpPointsDB,
		`SELECT id, name, phone_number, region, city, street, house_num FROM pick_up_points`)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	pickUpPoints := make([]model.PickUpPoint, len(pickUpPointsDB))
	for i := range pickUpPointsDB {
		pickUpPoints[i] = pickUpPointsDB[i].ConvertToPickUpPoint()
	}
	return pickUpPoints, nil
}
