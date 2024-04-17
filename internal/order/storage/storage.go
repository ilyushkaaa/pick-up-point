package storage

import (
	"context"

	"homework/internal/order/model"
)

//go:generate mockgen -source ./storage.go -destination=./mocks/storage.go -package=mock_storage
type OrderStorage interface {
	AddOrder(ctx context.Context, newOrder model.Order) error
	DeleteOrder(ctx context.Context, orderID uint64) error
	DeleteOrdersByPPID(ctx context.Context, ppID uint64) error
	IssueOrders(ctx context.Context, orderIDs map[uint64]struct{}) error
	GetUserOrders(ctx context.Context, clientID uint64) ([]model.Order, error)
	ReturnOrder(ctx context.Context, clientID, orderID uint64) error
	GetOrderReturns(ctx context.Context) ([]model.Order, error)
	GetOrderByID(ctx context.Context, ID uint64) (*model.Order, error)
	GetOrders(ctx context.Context) ([]model.Order, error)
}
