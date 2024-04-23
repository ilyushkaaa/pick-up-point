package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/filters/model"
	"homework/internal/order/delivery/grpc/dto"
	pb "homework/internal/pb/order"
	"homework/pkg/response"
)

func (o OrderDelivery) GetUserOrders(request *pb.GetUserOrdersRequest, server pb.Orders_GetUserOrdersServer) error {
	filters := model.Filters{
		Limit:  int(request.GetLimit()),
		PPOnly: request.PpOnly,
	}
	orders, err := o.service.GetUserOrders(server.Context(), request.GetClientId(), filters)
	if err != nil {
		o.logger.Errorf("internal server error in getting user orders: %v", err)
		return status.Errorf(codes.Internal, response.ErrInternal.Error())
	}

	for _, order := range orders {
		err = server.Send(dto.GetPBFromOrder(&order))
		if err != nil {
			o.logger.Errorf("error in sending order into stream: %v", err)
			return status.Errorf(codes.Internal, response.ErrInternal.Error())
		}
	}
	return nil
}
