package storage

import "homework/internal/pick-up_point/model"

func (fs *FilePPStorage) UpdatePickUpPoint(point model.PickUpPoint) error {
	fs.mu.RLock()
	for i, pp := range fs.cash {
		if pp.Name == point.Name {
			fs.mu.RUnlock()
			fs.mu.Lock()
			fs.cash[i] = point
			fs.mu.Unlock()
			return nil
		}
	}
	fs.mu.RUnlock()
	return nil
}
