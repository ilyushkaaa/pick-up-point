package service

import (
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage"
)

type PickUpPointService interface {
	AddPickUpPoint(point model.PickUpPoint) error
	GetPickUpPoints() ([]model.PickUpPoint, error)
	GetPickUpPointByName(name string) (*model.PickUpPoint, error)
	UpdatePickUpPoint(point model.PickUpPoint) error
}

type PPService struct {
	storage storage.PPStorage
}

func New(storage storage.PPStorage) *PPService {
	return &PPService{
		storage: storage,
	}
}
