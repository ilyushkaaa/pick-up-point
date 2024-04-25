package delivery

import (
	"context"
	"errors"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "homework/internal/pb/pick-up_point"
	"homework/internal/pick-up_point/storage"
	"homework/pkg/response"
)

func (P PPDelivery) Delete(ctx context.Context, req *pb.DeletePPRequest) (*pb.DeleteResponse, error) {
	ctx, span := P.tracer.Start(ctx, "DeletePickUpPoint")
	defer span.End()
	err := P.service.DeletePickUpPoint(ctx, req.Id)
	if err != nil {
		if errors.Is(err, storage.ErrPickUpPointNotFound) {
			P.logger.Errorf("no pick-up points with id %d", req.GetId())
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		P.logger.Errorf("internal server error in deleting pick-up point: %v", err)
		return nil, status.Errorf(codes.Internal, response.ErrInternal.Error())
	}
	P.cache.GoDeleteFromCache(ctx, strconv.FormatUint(req.GetId(), 10))
	return &pb.DeleteResponse{
		Result: "success",
	}, nil
}
