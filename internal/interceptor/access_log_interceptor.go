package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"homework/internal/events/model"
)

func (i *Interceptor) AccessLog(ctx context.Context, info *grpc.UnaryServerInfo) {
	p, _ := peer.FromContext(ctx)
	newEvent := model.NewEvent(p.Addr.String(), info.FullMethod)
	sendResult, err := i.producer.SendMessage(newEvent)
	if err != nil {
		i.logger.Errorf("error in writing new event into kafka: %s", err)
	} else {
		i.logger.Infof("message was sent to kafka: partition: %d, offset: %d", sendResult.Partition, sendResult.Offset)
	}
}
