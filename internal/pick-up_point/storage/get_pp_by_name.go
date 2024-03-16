package storage

import "homework/internal/pick-up_point/model"

func (fs *FilePPStorage) GetPickUpPointByName(name string) (*model.PickUpPoint, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	for _, pp := range fs.cache {
		if pp.Name == name {
			return &pp, nil
		}
	}
	return nil, nil
}
