package grpc

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/order/delivery/grpc/dto"
	"homework/internal/order/service"
	pb "homework/internal/pb/order"
	"homework/pkg/response"
)

func (o OrderDelivery) GetOrderReturns(request *pb.GetOrdersReturnsRequest, server pb.Orders_GetOrderReturnsServer) error {
	orders, err := o.service.GetOrderReturns(server.Context(), request.GetOrdersPerPage(), request.GetPageNum())
	if err != nil {
		if errors.Is(err, service.ErrNoOrdersOnThisPage) {
			o.logger.Errorf("no orders on page %d when %d orders per page", request.GetPageNum(), request.GetOrdersPerPage())
			return status.Errorf(codes.InvalidArgument, err.Error())
		}
		o.logger.Errorf("internal server error in getting order returns: %v", err)
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
