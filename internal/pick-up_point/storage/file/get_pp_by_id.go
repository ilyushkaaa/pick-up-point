package storage

import (
	"context"

	"homework/internal/pick-up_point/model"
)

func (fs *FilePPStorage) GetPickUpPointByID(_ context.Context, ID uint64) (*model.PickUpPoint, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	for _, pp := range fs.cache {
		if pp.ID == ID {
			return &pp, nil
		}
	}
	return nil, nil
}
