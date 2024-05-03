package delivery

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "homework/internal/pb/pick-up_point"
	"homework/internal/pick-up_point/delivery/grpc/dto"
	"homework/pkg/response"
)

func (P PPDelivery) GetAll(_ *pb.GetAllRequest, req pb.PickUpPoints_GetAllServer) error {
	ctx, span := P.tracer.Start(req.Context(), "GetAllPickUpPoints")
	defer span.End()
	pickUpPoints, err := P.service.GetPickUpPoints(ctx)
	if err != nil {
		P.logger.Errorf("internal server error in getting pick-up points: %v", err)
		return status.Errorf(codes.Internal, response.ErrInternal.Error())
	}
	for _, pp := range pickUpPoints {
		err = req.Send(dto.GetPBFromPickUpPoint(&pp))
		if err != nil {
			P.logger.Errorf("error in sending pick-up point into stream: %v", err)
			return status.Errorf(codes.Internal, response.ErrInternal.Error())
		}
	}
	return nil
}
