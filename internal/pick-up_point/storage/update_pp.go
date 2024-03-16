package storage

import "homework/internal/pick-up_point/model"

func (fs *FilePPStorage) UpdatePickUpPoint(point model.PickUpPoint) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	for i, pp := range fs.cache {
		if pp.Name == point.Name {
			fs.cache[i] = point
			return nil
		}
	}
	return nil
}
