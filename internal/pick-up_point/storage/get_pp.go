package storage

import "homework/internal/pick-up_point/model"

func (fs *FilePPStorage) GetPickUpPoints() ([]model.PickUpPoint, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	return fs.cache, nil
}
