package producer

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func newSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	syncProducerConfig := sarama.NewConfig()

	syncProducerConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	syncProducerConfig.Producer.RequiredAcks = sarama.WaitForAll
	syncProducerConfig.Producer.CompressionLevel = sarama.CompressionLevelDefault
	syncProducerConfig.Producer.Return.Successes = true
	syncProducerConfig.Producer.Return.Errors = true
	syncProducerConfig.Producer.Compression = sarama.CompressionGZIP

	syncProducer, err := sarama.NewSyncProducer(brokers, syncProducerConfig)
	if err != nil {
		return nil, fmt.Errorf("error in creating sync producer: %w", err)
	}

	return syncProducer, nil
}

func NewProducer(brokers []string) (*Producer, error) {
	syncProducer, err := newSyncProducer(brokers)
	if err != nil {
		return nil, err
	}

	producer := &Producer{
		brokers:      brokers,
		syncProducer: syncProducer,
	}

	return producer, nil
}
