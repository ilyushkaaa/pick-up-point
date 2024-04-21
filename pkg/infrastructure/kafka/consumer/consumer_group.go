package consumer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	"homework/internal/events/service/consumer"
	"homework/pkg/infrastructure/kafka"
)

const consumerStartTimeout = time.Second * 60

func Run(ctx context.Context, cfg *kafka.ConfigKafka, logger *zap.SugaredLogger, waitChan chan struct{}) error {
	client, err := newConsumerGroup(cfg)
	if err != nil {
		return err
	}

	eventService := consumer.NewEventsConsumer(logger)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err = client.Consume(ctx, []string{cfg.Topic}, &eventService); err != nil {
				logger.Errorf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-eventService.Ready()
	logger.Info("Sarama consumer up and running!...")
	waitChan <- struct{}{}

	wg.Wait()

	if err = client.Close(); err != nil {
		return err
	}
	return nil
}

func GoRunConsumer(ctx context.Context, cfg *kafka.ConfigKafka, logger *zap.SugaredLogger, waitChan chan struct{}) {
	go func() {
		err := Run(ctx, cfg, logger, waitChan)
		if err != nil {
			logger.Errorf("error in consumer running")
		}
	}()
}

func WaitForConsumerReady(waitChan chan struct{}) error {
	select {
	case <-time.After(consumerStartTimeout):
		return fmt.Errorf("timout for consumer start working")
	case <-waitChan:
		return nil
	}
}
