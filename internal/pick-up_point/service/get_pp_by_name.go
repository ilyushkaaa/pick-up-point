package service

import "homework/internal/pick-up_point/model"

func (ps *PPService) GetPickUpPointByNameService(name string) (*model.PickUpPoint, error) {
	pickUpPoint, err := ps.storage.GetPickUpPointByNameStorage(name)
	if err != nil {
		return nil, err
	}
	if pickUpPoint == nil {
		return nil, ErrPickUpPointNotFound
	}
	return pickUpPoint, nil
}
