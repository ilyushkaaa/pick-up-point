package storage

import (
	"os"

	"homework/Homework-1/internal/order/model"
)

type OrderStorage interface {
	AddOrderStorage(newOrder model.Order) error
	DeleteOrderStorage(orderID int) error
	IssueOrdersStorage(orderIDs map[int]struct{}) error
	GetUserOrdersStorage(clientID int) ([]model.Order, error)
	ReturnOrderStorage(clientID, orderID int) error
	GetOrderReturnsStorage() ([]model.Order, error)
	GetOrders() ([]model.Order, error)
}

type FileOrderStorage struct {
	file *os.File
}

func NewFileOrderStorage(file *os.File) *FileOrderStorage {
	return &FileOrderStorage{
		file: file,
	}
}
