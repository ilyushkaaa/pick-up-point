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

func (o OrderDelivery) Delete(ctx context.Context, request *pb.DeleteOrderRequest) (*pb.ResultResponse, error) {
	ctx, span := o.tracer.Start(ctx, "DeleteOrder")
	defer span.End()
	err := o.service.DeleteOrder(ctx, request.GetId())
	if err != nil {
		if errors.Is(err, storage.ErrOrderNotFound) {
			o.logger.Errorf("no orders with id %d", request.GetId())
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		if errors.Is(err, service.ErrOrderAlreadyIssued) {
			o.logger.Errorf("order with id %d is issued to user and is not in pick-up point", request.GetId())
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		if errors.Is(err, service.ErrOrderShelfLifeNotExpired) {
			o.logger.Errorf("shelf life for order %d is not expired", request.GetId())
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}

		o.logger.Errorf("internal server error in deleting order: %v", err)
		return nil, status.Errorf(codes.Internal, response.ErrInternal.Error())
	}
	return &pb.ResultResponse{Result: "success"}, nil
}
