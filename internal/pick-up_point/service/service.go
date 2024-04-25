package service

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"homework/internal/cache"
	orderStorage "homework/internal/order/storage"
	"homework/internal/pick-up_point/model"
	ppStorage "homework/internal/pick-up_point/storage"
	"homework/pkg/infrastructure/database/postgres/transaction_manager"
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
	cache              cache.Cache
	tracer             trace.Tracer
}

func New(ppStorage ppStorage.PPStorage, orderStorage orderStorage.OrderStorage,
	transactionManager transaction_manager.TransactionManager, cache cache.Cache, tracer trace.Tracer) *PPService {
	return &PPService{
		ppStorage:          ppStorage,
		orderStorage:       orderStorage,
		transactionManager: transactionManager,
		cache:              cache,
		tracer:             tracer,
	}
}
