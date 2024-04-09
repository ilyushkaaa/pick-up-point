package consumer

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"homework/internal/events/service"
)

func Run(brokers []string, logger *zap.SugaredLogger) error {
	client, err := newConsumerGroup(brokers)
	if err != nil {
		return err
	}

	keepRunning := true

	ctx, cancel := context.WithCancel(context.Background())

	eventService := service.NewEventsService(logger)
	consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, []string{"events"}, &eventService); err != nil {
				logger.Errorf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-eventService.Ready()
	logger.Info("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			logger.Info("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			logger.Info("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(client, &consumptionIsPaused, logger)
		}
	}

	cancel()
	wg.Wait()

	if err = client.Close(); err != nil {
		return err
	}
	return nil
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool, logger *zap.SugaredLogger) {
	if *isPaused {
		client.ResumeAll()
		logger.Info("Resuming consumption")
	} else {
		client.PauseAll()
		logger.Info("Pausing consumption")
	}

	*isPaused = !*isPaused
}
