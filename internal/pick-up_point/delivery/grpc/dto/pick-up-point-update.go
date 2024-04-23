package dto

import (
	pb "homework/internal/pb/pick-up_point"
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/model/address"
)

func GetPickUpPointFromPBUpdate(pp *pb.PickUpPointUpdate) model.PickUpPoint {
	return model.PickUpPoint{
		Name: pp.Name,
		Address: address.PPAddress{
			Region:   pp.Address.Region,
			City:     pp.Address.City,
			Street:   pp.Address.Street,
			HouseNum: pp.Address.HouseNum,
		},
		PhoneNumber: pp.PhoneNumber,
	}
}
