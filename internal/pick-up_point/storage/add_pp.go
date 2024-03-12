package storage

import (
	"homework/internal/pick-up_point/model"
)

func (fs *FilePPStorage) AddPickUpPoint(point model.PickUpPoint) error {
	fs.mu.Lock()
	fs.cash = append(fs.cash, point)
	return nil
}
