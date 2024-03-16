package service

import "homework/internal/pick-up_point/model"

func (ps *PPService) GetPickUpPointByName(name string) (*model.PickUpPoint, error) {
	pickUpPoint, err := ps.storage.GetPickUpPointByName(name)
	if err != nil {
		return nil, err
	}
	if pickUpPoint == nil {
		return nil, ErrPickUpPointNotFound
	}
	return pickUpPoint, nil
}
