package service

import "homework/internal/pick-up_point/model"

func (ps *PPService) AddPickUpPointService(point model.PickUpPoint) error {
	pickUpPoints, err := ps.storage.GetPickUpPointsStorage()
	if err != nil {
		return err
	}
	for _, pp := range pickUpPoints {
		if pp.Name == point.Name {
			return ErrPickUpPointAlreadyExists
		}
	}
	return ps.storage.AddPickUpPointStorage(point)
}
