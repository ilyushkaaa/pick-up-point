package storage

import "homework/internal/pick-up_point/model"

func (fs *FilePPStorage) GetPickUpPointByName(name string) (*model.PickUpPoint, error) {
	pickUpPoints, err := fs.GetPickUpPoints()
	if err != nil {
		return nil, err
	}
	for _, pp := range pickUpPoints {
		if pp.Name == name {
			return &pp, nil
		}
	}
	return nil, nil
}
