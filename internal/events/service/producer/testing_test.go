package producer

import (
	"testing"

	"github.com/golang/mock/gomock"
	producerMock "homework/pkg/kafka/producer/mocks"
)

type producerFixtures struct {
	ctrl          *gomock.Controller
	eventProducer *EventsProducer
	mockProducer  *producerMock.MockProducer
}

func setUp(t *testing.T) producerFixtures {
	ctrl := gomock.NewController(t)
	mockProducer := producerMock.NewMockProducer(ctrl)
	eventProducer := NewEventsProducer(mockProducer, "events")
	return producerFixtures{
		ctrl:          ctrl,
		mockProducer:  mockProducer,
		eventProducer: eventProducer,
	}
}
