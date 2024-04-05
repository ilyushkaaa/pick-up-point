package service

import (
	"context"

	filtermodel "homework/internal/filters/model"
	ordermodel "homework/internal/order/model"
	"homework/internal/order/service/packages"
	"homework/internal/order/storage"
)

type OrderService interface {
	AddOrder(ctx context.Context, order ordermodel.Order) error
	DeleteOrder(ctx context.Context, orderID uint64) error
	IssueOrders(ctx context.Context, orderIDs map[uint64]struct{}) error
	GetUserOrders(ctx context.Context, clientID uint64, filters filtermodel.Filters) ([]ordermodel.Order, error)
	ReturnOrder(ctx context.Context, clientID, orderID uint64) error
	GetOrderReturns(ctx context.Context, maxOrdersPerPage, pageNum uint64) ([]ordermodel.Order, error)
}

// PP - pick-up point

type OrderServicePP struct {
	storage  storage.OrderStorage
	packages map[string]*packages.Package
}

func New(storage storage.OrderStorage, packages map[string]*packages.Package) *OrderServicePP {
	return &OrderServicePP{
		storage:  storage,
		packages: packages,
	}
}
