package storage

import (
	"homework/internal/pick-up_point/model"
)

func (fs *FilePPStorage) AddPickUpPoint(point model.PickUpPoint) error {
	pickUpPoints, err := fs.GetPickUpPoints()
	if err != nil {
		return err
	}
	pickUpPoints = append(pickUpPoints, point)
	return fs.writePickUpPoints(pickUpPoints)
}
