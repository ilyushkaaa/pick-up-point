//go:build integration
// +build integration

package events

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/require"
)

type helper struct {
	waitChan chan struct{}
}

func (h helper) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h helper) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h helper) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		h.waitChan <- struct{}{}
		session.MarkMessage(msg, "")
	}
	return nil
}

func HelperRun(handler helper, brokers []string, ctx context.Context) error {

	config2 := sarama.NewConfig()
	config2.Consumer.Group.Session.Timeout = 10 * time.Second
	config2.Consumer.Group.Heartbeat.Interval = 3 * time.Second
	config2.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup2, err := sarama.NewConsumerGroup(brokers, "group-2", config2)
	if err != nil {
		fmt.Printf("Error creating consumer group 2: %v\n", err)
		return err
	}
	defer consumerGroup2.Close()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			err = consumerGroup2.Consume(ctx, []string{eventsTopic}, handler)
			if err != nil {
				fmt.Printf("Error consuming messages with consumer group 2: %v\n", err)
				return
			}
		}
	}()

	wg.Wait()
	return nil
}

func GoHelperRun(t *testing.T, handler helper, brokers []string, ctx context.Context) {
	t.Helper()

	go func() {
		err := HelperRun(handler, brokers, ctx)
		require.NoError(t, err)
	}()

}
