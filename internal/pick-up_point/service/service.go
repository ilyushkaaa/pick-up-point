package service

import (
	"context"

	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage"
)
//go:generate mockgen -source ./service.go -destination=./mocks/service.go -package=mock_service
type PickUpPointService interface {
	AddPickUpPoint(ctx context.Context, point model.PickUpPoint) (*model.PickUpPoint, error)
	GetPickUpPoints(ctx context.Context) ([]model.PickUpPoint, error)
	GetPickUpPointByID(ctx context.Context, ID uint64) (*model.PickUpPoint, error)
	UpdatePickUpPoint(ctx context.Context, point model.PickUpPoint) error
	DeletePickUpPoint(ctx context.Context, ID uint64) error
}

type PPService struct {
	storage storage.PPStorage
}

func New(storage storage.PPStorage) *PPService {
	return &PPService{
		storage: storage,
	}
}
