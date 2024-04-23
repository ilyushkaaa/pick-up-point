package dto

import (
	pb "homework/internal/pb/pick-up_point"
	"homework/internal/pick-up_point/model"
	"homework/internal/pick-up_point/model/address"
)

func GetPickUpPointFromPBAdd(pp *pb.PickUpPointAdd) model.PickUpPoint {
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

func GetPBFromPickUpPoint(pp *model.PickUpPoint) *pb.PickUpPoint {
	return &pb.PickUpPoint{
		Id:   pp.ID,
		Name: pp.Name,
		Address: &pb.Address{
			Region:   pp.Address.Region,
			City:     pp.Address.City,
			Street:   pp.Address.Street,
			HouseNum: pp.Address.HouseNum,
		},
		PhoneNumber: pp.PhoneNumber,
	}
}
