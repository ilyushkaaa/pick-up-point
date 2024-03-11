package storage

import (
	"encoding/json"

	"homework/internal/pick-up_point/model"
)

func (fs *FilePPStorage) writePickUpPoints(pickUpPoints []model.PickUpPoint) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	err := fs.file.Truncate(0)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(fs.file)
	return encoder.Encode(pickUpPoints)
}
