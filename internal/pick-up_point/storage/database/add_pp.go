package storage

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (s *PPStorageDB) AddPickUpPoint(ctx context.Context, point model.PickUpPoint) (*model.PickUpPoint, error) {
	var id uint64
	err := s.db.Cluster.QueryRow(ctx,
		`INSERT INTO pick_up_points (name, phone_number, region, city, street, house_num)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		point.Name, point.PhoneNumber, point.Address.Region, point.Address.City, point.Address.Street, point.Address.HouseNum).Scan(&id)
	if err != nil {
		return nil, err
	}
	point.ID = id
	return &point, nil
}
