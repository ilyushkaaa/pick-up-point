package delivery

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "homework/internal/pb/pick-up_point"
	"homework/internal/pick-up_point/delivery/grpc/dto"
	"homework/internal/pick-up_point/service"
	"homework/pkg/response"
)

func (P PPDelivery) Add(ctx context.Context, req *pb.PickUpPointAdd) (*pb.PickUpPoint, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	addedPP, err := P.service.AddPickUpPoint(ctx, dto.GetPickUpPointFromPBAdd(req))
	if err != nil {
		if errors.Is(err, service.ErrPickUpPointAlreadyExists) {
			P.logger.Errorf("pick-up point with name %s already exists", req.Name)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		P.logger.Errorf("internal server error in adding pick-up point: %v", err)
		return nil, status.Errorf(codes.Internal, response.ErrInternal.Error())
	}
	return dto.GetPBFromPickUpPoint(addedPP), nil
}
