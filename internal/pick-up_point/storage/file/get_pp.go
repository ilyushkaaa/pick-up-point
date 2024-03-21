package storage

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (fs *FilePPStorage) GetPickUpPoints(_ context.Context) ([]model.PickUpPoint, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	return fs.cache, nil
}
