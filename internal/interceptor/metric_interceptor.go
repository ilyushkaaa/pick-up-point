package interceptor

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func (i *Interceptor) Metric(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	i.metrics.Hits.Inc()
	startTime := time.Now()

	statusCode := "0"
	res, err := handler(ctx, req)
	endTime := time.Now()
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			statusCode = statusErr.Code().String()
		}
	}

	i.metrics.HitsByResponseCodes.WithLabelValues(statusCode, info.FullMethod).Inc()
	i.metrics.HitsByResponseCodesAndRequestTime.WithLabelValues(statusCode,
		info.FullMethod).Observe(endTime.Sub(startTime).Seconds())
	return res, err
}
