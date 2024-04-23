package interceptor

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"homework/internal/events/service/producer"
)

type Interceptor struct {
	logger   *zap.SugaredLogger
	producer *producer.EventsProducer
}

func New(logger *zap.SugaredLogger, producer *producer.EventsProducer) *Interceptor {
	return &Interceptor{
		logger:   logger,
		producer: producer,
	}
}

func (i *Interceptor) CallInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	i.AccessLog(ctx, info)
	return i.Auth(ctx, req, handler)
}
