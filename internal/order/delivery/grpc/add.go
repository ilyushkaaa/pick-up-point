package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/order/delivery/grpc/dto"
	"homework/internal/order/service"
	pb "homework/internal/pb/order"
	"homework/internal/pick-up_point/storage"
	"homework/pkg/response"
)

func (o OrderDelivery) Add(ctx context.Context, data *pb.OrderFromCourierInputData) (*pb.OrderFromCourierInputData, error) {
	err := o.service.AddOrder(ctx, dto.GetOrderFromPB(data))
	if err != nil {
		if errors.Is(err, service.ErrOrderAlreadyInPickUpPoint) {
			o.logger.Errorf("order with id %d already in pick-up point", data.GetId())
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		if errors.Is(err, storage.ErrPickUpPointNotFound) {
			o.logger.Errorf("no pick-up points with id %d", data.GetPickUpPointId())
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		if errors.Is(err, service.ErrShelfTimeExpired) {
			o.logger.Errorf("shelf time for this order has expired: %v", data.GetStorageExpirationDate())
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		if errors.Is(err, service.ErrUnknownPackage) {
			o.logger.Errorf("unknown package type %s", data.PackageType)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		if errors.Is(err, service.ErrPackageCanNotBeApplied) {
			o.logger.Errorf("%s can not be applied for order %v", data.PackageType, data)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		o.logger.Errorf("internal server error in adding order: %v", err)
		return nil, status.Errorf(codes.Internal, response.ErrInternal.Error())
	}
	return data, nil
}
