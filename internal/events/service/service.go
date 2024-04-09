package service

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"homework/internal/events/model"
)

type EventsService struct {
	logger *zap.SugaredLogger
	ready  chan bool
}

func NewEventsService(logger *zap.SugaredLogger) EventsService {
	return EventsService{
		logger: logger,
		ready:  make(chan bool),
	}
}

func (s *EventsService) Ready() <-chan bool {
	return s.ready
}

func (s *EventsService) Setup(_ sarama.ConsumerGroupSession) error {
	close(s.ready)

	return nil
}

func (s *EventsService) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (s *EventsService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():

			newEvent := &model.Event{}
			err := json.Unmarshal(message.Value, &newEvent)
			if err != nil {
				s.logger.Errorf("consumer group error: %v", err)
			}

			s.logger.Infow("New request",
				"request time", newEvent.RequestTime,
				"method", newEvent.RequestMethod,
				"remote_addr", newEvent.RemoteAddr,
				"url", newEvent.URL,
				"request body", newEvent.RequestBody,
			)

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
