package storage

import (
	"homework/internal/pick-up_point/model"
)

func (fs *FilePPStorage) AddPickUpPoint(point model.PickUpPoint) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	fs.cache = append(fs.cache, point)
	return nil
}
