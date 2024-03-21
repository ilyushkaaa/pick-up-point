package storage

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (fs *FilePPStorage) UpdatePickUpPoint(_ context.Context, point model.PickUpPoint) (bool, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	for i, pp := range fs.cache {
		if pp.Name == point.Name {
			fs.cache[i] = point
			return true, nil
		}
	}
	return false, nil
}
