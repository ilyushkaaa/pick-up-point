package storage

import (
	"fmt"
	"os"
	"sync"

	"homework/internal/pick-up_point/model"
)

const fileName = "storage_pp.json"

type PPStorage interface {
	AddPickUpPoint(point model.PickUpPoint) error
	GetPickUpPoints() ([]model.PickUpPoint, error)
	GetPickUpPointByName(name string) (*model.PickUpPoint, error)
	UpdatePickUpPoint(point model.PickUpPoint) error
	Close() error
}

type FilePPStorage struct {
	file *os.File
	mu   *sync.Mutex
}

func New() (*FilePPStorage, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("error in opening file")
	}
	return &FilePPStorage{
		file: file,
		mu:   &sync.Mutex{},
	}, nil
}
