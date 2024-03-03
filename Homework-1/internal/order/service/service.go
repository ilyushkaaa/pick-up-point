package service

import (
	"time"

	"homework/Homework-1/internal/order/model"
	"homework/Homework-1/internal/order/storage"
)

type OrderService interface {
	AddOrderService(orderID, clientID int, expireDate time.Time) error
	DeleteOrderService(orderID int) error
	IssueOrderService(orderIDs map[int]struct{}) error
	GetUserOrdersService(clientID, limit int, ppOnly bool) ([]model.Order, error)
	ReturnOrderService(clientID, orderID int) error
	GetOrderReturnsService() ([][]model.Order, error)
}

// PP - pick-up point

type OrderServicePP struct {
	storage storage.OrderStorage
}

func NewOrderServicePP(storage storage.OrderStorage) *OrderServicePP {
	return &OrderServicePP{
		storage: storage,
	}
}
