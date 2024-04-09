package producer

import (
	"github.com/Shopify/sarama"
)

type Producer struct {
	brokers      []string
	syncProducer sarama.SyncProducer
}

func (k *Producer) SendSyncMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return k.syncProducer.SendMessage(message)
}

func (k *Producer) Close() error {
	return k.syncProducer.Close()
}
