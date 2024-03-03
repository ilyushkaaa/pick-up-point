package storage

import (
	"encoding/json"

	"homework/Homework-1/internal/order/model"
)

func (fs *FileOrderStorage) writeOrders(orders []model.Order) error {
	err := fs.file.Truncate(0)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(fs.file)
	return encoder.Encode(orders)
}
