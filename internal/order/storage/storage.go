package storage

import (
	"fmt"
	"os"

	"homework/internal/order/model"
)

const fileName = "storage.json"

type OrderStorage interface {
	AddOrderStorage(newOrder model.Order) error
	DeleteOrderStorage(orderID int) error
	IssueOrdersStorage(orderIDs map[int]struct{}) error
	GetUserOrdersStorage(clientID int) ([]model.Order, error)
	ReturnOrderStorage(clientID, orderID int) error
	GetOrderReturnsStorage() ([]model.Order, error)
	GetOrders() ([]model.Order, error)
	Close() error
}

type FileOrderStorage struct {
	file *os.File
}

func New() (*FileOrderStorage, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("error in opening file")
	}
	return &FileOrderStorage{
		file: file,
	}, nil
}
