package storage

import (
	"encoding/json"
	"errors"
	"io"

	"homework/internal/pick-up_point/model"
)

func (fs *FilePPStorage) GetPickUpPointsStorage() ([]model.PickUpPoint, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	decoder := json.NewDecoder(fs.file)
	var pickUpPoints []model.PickUpPoint

	if err := decoder.Decode(&pickUpPoints); err != nil {
		if !errors.Is(err, io.EOF) {
			return nil, err
		}
	}
	_, err := fs.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return pickUpPoints, nil

}
