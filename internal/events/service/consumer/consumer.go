package consumer

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"homework/internal/events/model"
)

type EventsConsumer struct {
	logger *zap.SugaredLogger
	ready  chan bool
}

func NewEventsConsumer(logger *zap.SugaredLogger) EventsConsumer {
	return EventsConsumer{
		logger: logger,
		ready:  make(chan bool),
	}
}

func (s *EventsConsumer) Ready() <-chan bool {
	return s.ready
}

func (s *EventsConsumer) Setup(_ sarama.ConsumerGroupSession) error {
	close(s.ready)

	return nil
}

func (s *EventsConsumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (s *EventsConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():

			newEvent := &model.Event{}
			err := json.Unmarshal(message.Value, &newEvent)
			if err != nil {
				s.logger.Errorf("can not unmarshal message: %v", err)
				continue
			}

			s.logger.Infow("New request",
				"request time", newEvent.RequestTime,
				"method", newEvent.RequestMethod,
				"remote_addr", newEvent.RemoteAddr,
				"url", newEvent.URL,
			)

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
