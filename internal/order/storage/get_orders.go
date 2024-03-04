package storage

import (
	"encoding/json"
	"errors"
	"io"

	"homework/internal/order/model"
)

func (fs *FileOrderStorage) GetOrders() ([]model.Order, error) {
	decoder := json.NewDecoder(fs.file)
	var orders []model.Order

	if err := decoder.Decode(&orders); err != nil {
		if !errors.Is(err, io.EOF) {
			return nil, err
		}
	}
	_, err := fs.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
