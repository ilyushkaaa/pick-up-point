package service

import "homework/internal/pick-up_point/model"

func (ps *PPService) GetPickUpPoints() ([]model.PickUpPoint, error) {
	pickUpPoints, err := ps.storage.GetPickUpPoints()
	if err != nil {
		return nil, err
	}
	if len(pickUpPoints) == 0 {
		return nil, ErrNoPickUpPoints
	}
	return pickUpPoints, nil
}
