package delivery

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "homework/internal/pb/pick-up_point"
	"homework/internal/pick-up_point/delivery/grpc/dto"
	"homework/internal/pick-up_point/storage"
	"homework/pkg/response"
)

func (P PPDelivery) Update(ctx context.Context, req *pb.PickUpPointUpdate) (*pb.PickUpPointUpdate, error) {
	ctx, span := P.tracer.Start(ctx, "UpdatePickUpPoint")
	defer span.End()
	err := req.ValidateAll()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	err = P.service.UpdatePickUpPoint(ctx, dto.GetPickUpPointFromPBUpdate(req))
	if err != nil {
		if errors.Is(err, storage.ErrPickUpPointNotFound) {
			P.logger.Errorf("pick-up point with id %d does not exist", req.GetId())
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		if errors.Is(err, storage.ErrPickUpPointNameExists) {
			P.logger.Errorf("pick-up point with name %s already exists", req.GetName())
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		P.logger.Errorf("internal server error in updating pick-up point: %v", err)
		return nil, status.Errorf(codes.Internal, response.ErrInternal.Error())
	}
	return req, nil
}
