package producer

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"homework/internal/events/model"
	"homework/pkg/kafka/producer"
)

type EventsProducer struct {
	logger   *zap.SugaredLogger
	producer producer.Producer
	topic    string
}

func NewEventsProducer(producer producer.Producer, topic string, logger *zap.SugaredLogger) *EventsProducer {
	return &EventsProducer{
		producer: producer,
		topic:    topic,
		logger:   logger,
	}
}

func (s *EventsProducer) SendMessage(event model.Event) error {
	kafkaMsg, err := s.BuildMessage(event)
	if err != nil {
		s.logger.Errorf("send message marshal error: %s", err)
		return err
	}

	partition, offset, err := s.producer.SendMessage(kafkaMsg)

	if err != nil {
		s.logger.Errorf("send message connector error: %s", err)
		return err
	}

	s.logger.Infof("message was sent to kafka: partition: %d, offset: %d", partition, offset)
	return nil
}

func (s *EventsProducer) BuildMessage(event model.Event) (*sarama.ProducerMessage, error) {
	msg, err := json.Marshal(event)

	if err != nil {
		s.logger.Infof("send message marshal error: %s", err)
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(msg),
		Partition: -1,
	}, nil
}
