package storage

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (fs *FilePPStorage) GetPickUpPointByID(_ context.Context, id uint64) (*model.PickUpPoint, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	for _, pp := range fs.cache {
		if pp.ID == id {
			return &pp, nil
		}
	}
	return nil, nil
}
