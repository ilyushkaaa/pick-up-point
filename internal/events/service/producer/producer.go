package producer

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"homework/internal/events/model"
	"homework/pkg/infrastructure/kafka/producer"
)

type EventsProducer struct {
	producer producer.Producer
	topic    string
}

type SendMessageResult struct {
	Partition int32
	Offset    int64
}

func New(producer producer.Producer, topic string) *EventsProducer {
	return &EventsProducer{
		producer: producer,
		topic:    topic,
	}
}

func (s *EventsProducer) SendMessage(event model.Event) (*SendMessageResult, error) {
	kafkaMsg, err := s.BuildMessage(event)
	if err != nil {
		return nil, err
	}

	partition, offset, err := s.producer.SendMessage(kafkaMsg)

	if err != nil {
		return nil, err
	}
	return &SendMessageResult{
		Partition: partition,
		Offset:    offset}, nil
}

func (s *EventsProducer) BuildMessage(event model.Event) (*sarama.ProducerMessage, error) {
	msg, err := json.Marshal(event)

	if err != nil {
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(msg),
		Partition: -1,
	}, nil
}

func (s *EventsProducer) Close() error {
	return s.producer.Close()
}
