package storage

import (
	"context"

	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage/database/dto"
)

func (s *PPStorageDB) AddPickUpPoint(ctx context.Context, point model.PickUpPoint) (*model.PickUpPoint, error) {
	var id uint64
	pointDB := dto.NewPickUpPointDB(point)
	err := s.db.Cluster.QueryRow(ctx,
		`INSERT INTO pick_up_points (name, phone_number, region, city, street, house_num)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		pointDB.Name, pointDB.PhoneNumber, pointDB.Region, pointDB.City, pointDB.Street, pointDB.HouseNum).Scan(&id)
	if err != nil {
		return nil, err
	}
	point.ID = id
	return &point, nil
}
