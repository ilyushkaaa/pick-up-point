package service

import (
	"context"

	"homework/internal/cache"
	filtermodel "homework/internal/filters/model"
	ordermodel "homework/internal/order/model"
	"homework/internal/order/service/packages"
	orderStorage "homework/internal/order/storage"
	ppStorage "homework/internal/pick-up_point/storage"
	"homework/pkg/infrastructure/database/postgres/transaction_manager"
)

type OrderService interface {
	AddOrder(ctx context.Context, order ordermodel.Order) error
	DeleteOrder(ctx context.Context, orderID uint64) error
	IssueOrders(ctx context.Context, orderIDs []uint64) error
	GetUserOrders(ctx context.Context, clientID uint64, filters filtermodel.Filters) ([]ordermodel.Order, error)
	ReturnOrder(ctx context.Context, clientID, orderID uint64) error
	GetOrderReturns(ctx context.Context, maxOrdersPerPage, pageNum uint64) ([]ordermodel.Order, error)
}

// PP - pick-up point

type OrderServicePP struct {
	orderStorage       orderStorage.OrderStorage
	ppStorage          ppStorage.PPStorage
	packages           map[string]*packages.Package
	transactionManager transaction_manager.TransactionManager
	cache              cache.Cache
}

func New(storage orderStorage.OrderStorage, ppStorage ppStorage.PPStorage, packages map[string]*packages.Package,
	transactionManager transaction_manager.TransactionManager, cache cache.Cache) *OrderServicePP {
	return &OrderServicePP{
		orderStorage:       storage,
		ppStorage:          ppStorage,
		packages:           packages,
		transactionManager: transactionManager,
		cache:              cache,
	}
}
