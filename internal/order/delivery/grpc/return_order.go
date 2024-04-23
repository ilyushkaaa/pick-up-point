package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/order/service"
	"homework/internal/order/storage"
	pb "homework/internal/pb/order"
	"homework/pkg/response"
)

func (o OrderDelivery) ReturnOrders(ctx context.Context, data *pb.ReturnOrderInputData) (*pb.ResultResponse, error) {
	err := o.service.ReturnOrder(ctx, data.GetClientId(), data.GetOrderId())
	if err != nil {
		if errors.Is(err, service.ErrClientOrderNotFound) || errors.Is(err, storage.ErrClientOrderNotFound) {
			o.logger.Errorf("client with id %d has not got order with id %d", data.GetClientId(), data.GetOrderId())
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		if errors.Is(err, service.ErrOrderIsNotIssued) {
			o.logger.Errorf("order with id %d is not issued", data.GetOrderId())
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		if errors.Is(err, service.ErrOrderIsReturned) {
			o.logger.Errorf("order with id %d is already returned", data.GetOrderId())
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		if errors.Is(err, service.ErrReturnTimeExpired) {
			o.logger.Errorf("return time for order %d has expired", data.GetOrderId())
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		o.logger.Errorf("internal server error in returning order: %v", err)
		return nil, status.Errorf(codes.Internal, response.ErrInternal.Error())
	}
	return &pb.ResultResponse{Result: "success"}, nil
}
