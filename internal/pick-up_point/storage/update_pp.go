package storage

import "homework/internal/pick-up_point/model"

func (fs *FilePPStorage) UpdatePickUpPoint(point model.PickUpPoint) error {
	pickUpPoints, err := fs.GetPickUpPoints()
	if err != nil {
		return err
	}
	for i, pp := range pickUpPoints {
		if pp.Name == point.Name {
			pickUpPoints[i] = point
			break
		}
	}
	return fs.writePickUpPoints(pickUpPoints)
}
