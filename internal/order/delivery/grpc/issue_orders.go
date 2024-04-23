package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/order/service"
	pb "homework/internal/pb/order"
	"homework/pkg/response"
)

func (o OrderDelivery) IssueOrders(ctx context.Context, issue *pb.OrdersToIssue) (*pb.ResultResponse, error) {
	err := o.service.IssueOrders(ctx, issue.OrderIds)
	if err != nil {
		if errors.Is(err, service.ErrOrdersOfDifferentClients) {
			o.logger.Error("passed orders belong to different clients")
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		if errors.Is(err, service.ErrOrderAlreadyIssued) {
			o.logger.Error("there are orders which already issued")
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		if errors.Is(err, service.ErrNotAllOrdersWereFound) {
			o.logger.Error("some of passed orders does not exist")
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		o.logger.Errorf("internal server error in issuing orders: %v", err)
		return nil, status.Errorf(codes.Internal, response.ErrInternal.Error())
	}
	return &pb.ResultResponse{Result: "success"}, nil
}
