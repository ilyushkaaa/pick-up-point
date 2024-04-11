package consumer

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"homework/internal/events/service/consumer"
)

func Run(brokers []string, logger *zap.SugaredLogger, ctx context.Context, topic, groupID string) error {
	client, err := newConsumerGroup(brokers, groupID)
	if err != nil {
		return err
	}

	eventService := consumer.NewEventsConsumer(logger)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err = client.Consume(ctx, []string{topic}, &eventService); err != nil {
				logger.Errorf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-eventService.Ready()
	logger.Info("Sarama consumer up and running!...")

	wg.Wait()

	if err = client.Close(); err != nil {
		return err
	}
	return nil
}
