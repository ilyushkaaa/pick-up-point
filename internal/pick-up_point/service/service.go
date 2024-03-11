package service

import (
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage"
)

type PickUpPointService interface {
	AddPickUpPointService(point model.PickUpPoint) error
	GetPickUpPointsService() ([]model.PickUpPoint, error)
	GetPickUpPointByNameService(name string) (*model.PickUpPoint, error)
}

type PPService struct {
	storage storage.PPStorage
}

func New(storage storage.PPStorage) *PPService {
	return &PPService{
		storage: storage,
	}
}
