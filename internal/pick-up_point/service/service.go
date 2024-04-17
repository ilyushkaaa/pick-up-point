package service

import (
	"context"

	orderStorage "homework/internal/order/storage"
	"homework/internal/pick-up_point/model"
	ppStorage "homework/internal/pick-up_point/storage"
	"homework/pkg/database/postgres/transaction_manager"
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
	orderStorage       orderStorage.OrderStorage
	ppStorage          ppStorage.PPStorage
	transactionManager transaction_manager.TransactionManager
}

func New(storage ppStorage.PPStorage, orderStorage orderStorage.OrderStorage,
	transactionManager transaction_manager.TransactionManager) *PPService {
	return &PPService{
		ppStorage:          storage,
		orderStorage:       orderStorage,
		transactionManager: transactionManager,
	}
}
