package kafka

import (
	"os"
	"strings"

	"github.com/Shopify/sarama"
	eventsProducer "homework/internal/events/service/producer"
	"homework/pkg/kafka/producer"
)

type ConfigKafka struct {
	Brokers         []string
	ConsumerGroupID string
	Topic           string
	Producer        *eventsProducer.EventsProducer
	ConsumerConfig  *sarama.Config
}

func NewConfig() (*ConfigKafka, error) {
	brokersFromEnv := os.Getenv("KAFKA_BROKERS")
	brokers := strings.Split(brokersFromEnv, ",")
	syncProducer, err := producer.New(brokers)
	if err != nil {
		return nil, err
	}
	topic := os.Getenv("KAFKA_EVENTS_TOPIC")
	groupID := os.Getenv("EVENTS_CONSUMER_GROUP_ID")
	producerEvents := eventsProducer.New(syncProducer, topic)
	return &ConfigKafka{
		Brokers:         brokers,
		ConsumerGroupID: groupID,
		Topic:           topic,
		Producer:        producerEvents,
	}, nil
}

func (c *ConfigKafka) Close() error {
	return c.Producer.Close()
}
