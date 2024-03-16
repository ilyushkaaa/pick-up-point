package service

import "homework/internal/pick-up_point/model"

func (ps *PPService) UpdatePickUpPoint(point model.PickUpPoint) error {
	pickUpPoints, err := ps.storage.GetPickUpPoints()
	if err != nil {
		return err
	}
	for _, pp := range pickUpPoints {
		if pp.Name == point.Name {
			return ps.storage.UpdatePickUpPoint(point)
		}
	}
	return ErrPickUpPointNotFound
}
