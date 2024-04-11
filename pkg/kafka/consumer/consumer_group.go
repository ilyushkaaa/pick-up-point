package consumer

import (
	"context"
	"os"
	"sync"

	"go.uber.org/zap"
	"homework/internal/events/service/consumer"
)

func Run(brokers []string, logger *zap.SugaredLogger, ctx context.Context) error {
	topic := os.Getenv("KAFKA_EVENTS_TOPIC")
	client, err := newConsumerGroup(brokers)
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
