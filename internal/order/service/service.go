package service

import (
	"time"

	filtermodel "homework/internal/filters/model"
	ordermodel "homework/internal/order/model"
	"homework/internal/order/storage"
)

type OrderService interface {
	AddOrderService(orderID, clientID int, expireDate time.Time) error
	DeleteOrderService(orderID int) error
	IssueOrderService(orderIDs map[int]struct{}) error
	GetUserOrdersService(clientID int, filters filtermodel.Filters) ([]ordermodel.Order, error)
	ReturnOrderService(clientID, orderID int) error
	GetOrderReturnsService(pageNum int) ([]ordermodel.Order, error)
}

// PP - pick-up point

type OrderServicePP struct {
	storage storage.OrderStorage
}

func New(storage storage.OrderStorage) *OrderServicePP {
	return &OrderServicePP{
		storage: storage,
	}
}
