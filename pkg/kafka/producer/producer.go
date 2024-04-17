package producer

import (
	"github.com/Shopify/sarama"
)

//go:generate mockgen -source ./producer.go -destination=./mocks/producer.go -package=producer
type Producer interface {
	SendMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error)
	Close() error
}
type SyncProducer struct {
	brokers      []string
	syncProducer sarama.SyncProducer
}

func (k *SyncProducer) SendMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return k.syncProducer.SendMessage(message)
}

func (k *SyncProducer) Close() error {
	return k.syncProducer.Close()
}
