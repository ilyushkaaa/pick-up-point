package storage

import (
	"context"

	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/storage"
)

func (fs *FilePPStorage) GetPickUpPointByName(_ context.Context, name string) (*model.PickUpPoint, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	for _, pp := range fs.cache {
		if pp.Name == name {
			return &pp, nil
		}
	}
	return nil, storage.ErrPickUpPointNotFound
}
