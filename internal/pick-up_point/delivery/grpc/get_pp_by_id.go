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

func (P PPDelivery) GetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.PickUpPoint, error) {
	ctx, span := P.tracer.Start(ctx, "GetPickUpPointByID")
	defer span.End()
	pp, err := P.service.GetPickUpPointByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, storage.ErrPickUpPointNotFound) {
			P.logger.Errorf("no pick-up points with id %d", req.GetId())
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		P.logger.Errorf("internal server error in getting pick-up point: %v", err)
		return nil, status.Errorf(codes.Internal, response.ErrInternal.Error())
	}
	return dto.GetPBFromPickUpPoint(pp), nil
}
