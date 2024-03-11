package storage

import "homework/internal/pick-up_point/model"

func (fs *FilePPStorage) GetPickUpPointByNameStorage(name string) (*model.PickUpPoint, error) {
	pickUpPoints, err := fs.GetPickUpPointsStorage()
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
