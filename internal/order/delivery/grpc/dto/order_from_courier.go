package dto

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"homework/internal/order/model"
	pb "homework/internal/pb/order"
)

func GetOrderFromPB(data *pb.OrderFromCourierInputData) model.Order {
	order := model.Order{
		ID:                    data.Id,
		ClientID:              data.ClientId,
		Weight:                data.Weight,
		Price:                 data.Price,
		StorageExpirationDate: data.StorageExpirationDate.AsTime(),
		PickUpPointID:         data.PickUpPointId,
	}
	if data.PackageType != nil {
		order.PackageType = *data.PackageType
	}
	return order
}

func GetPBFromOrder(order *model.Order) *pb.Order {
	return &pb.Order{
		Id:                    order.ID,
		ClientId:              order.ClientID,
		Weight:                order.Weight,
		Price:                 order.Price,
		PackageType:           order.PackageType,
		StorageExpirationDate: timestamppb.New(order.StorageExpirationDate),
		OrderIssueDate:        timestamppb.New(order.OrderIssueDate),
		IsIssued:              order.IsIssued,
		IsReturned:            order.IsReturned,
		PickUpPointId:         order.PickUpPointID,
	}
}
