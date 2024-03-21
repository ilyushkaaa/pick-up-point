package storage

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (fs *FilePPStorage) AddPickUpPoint(_ context.Context, point model.PickUpPoint) (*model.PickUpPoint, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	point.ID = fs.nextID
	fs.nextID++
	fs.cache = append(fs.cache, point)
	return &point, nil
}
