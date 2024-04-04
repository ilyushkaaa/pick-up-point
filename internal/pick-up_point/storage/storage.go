package storage

import (
	"context"

	"homework/internal/pick-up_point/model"
)

//go:generate mockgen -source ./storage.go -destination=./mocks/storage.go -package=mock_storage
type PPStorage interface {
	AddPickUpPoint(ctx context.Context, point model.PickUpPoint) (*model.PickUpPoint, error)
	GetPickUpPoints(ctx context.Context) ([]model.PickUpPoint, error)
	GetPickUpPointByID(ctx context.Context, ID uint64) (*model.PickUpPoint, error)
	GetPickUpPointByName(ctx context.Context, name string) (*model.PickUpPoint, error)
	UpdatePickUpPoint(ctx context.Context, point model.PickUpPoint) error
	DeletePickUpPoint(ctx context.Context, ID uint64) error
}
