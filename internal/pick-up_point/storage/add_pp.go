package storage

import (
	"encoding/json"

	"homework/internal/pick-up_point/model"
)

func (fs *FilePPStorage) AddPickUpPointStorage(point model.PickUpPoint) error {
	pickUpPoints, err := fs.GetPickUpPointsStorage()
	if err != nil {
		return err
	}
	pickUpPoints = append(pickUpPoints, point)
	fs.mu.Lock()
	defer fs.mu.Unlock()
	err = fs.file.Truncate(0)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(fs.file)
	return encoder.Encode(pickUpPoints)
}
