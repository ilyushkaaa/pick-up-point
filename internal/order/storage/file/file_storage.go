package storage

import (
	"fmt"
	"os"
)

const fileName = "storage_order.json"

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
