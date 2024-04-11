package producer

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"homework/tests/fixtures"
	"homework/tests/test_json"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()

	t.Run("unsuccessful send", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockProducer.EXPECT().SendMessage(gomock.Any()).Return(int32(1), int64(1), fmt.Errorf("error"))

		msg := s.eventProducer.SendMessage(fixtures.Event().Valid().Value())

		assert.Error(t, msg.Error)
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		expected := SendMessageResult{
			Partition: 1,
			Offset:    1,
			Error:     nil,
		}
		s.mockProducer.EXPECT().SendMessage(gomock.Any()).Return(int32(1), int64(1), nil)

		msg := s.eventProducer.SendMessage(fixtures.Event().Valid().Value())

		assert.Equal(t, expected, msg)
	})
}

func TestBuildMessage(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		expected := &sarama.ProducerMessage{
			Topic:     s.eventProducer.topic,
			Value:     sarama.ByteEncoder(test_json.ValidEventJSON),
			Partition: -1,
		}

		msg, err := s.eventProducer.BuildMessage(fixtures.Event().Valid().Value())

		assert.NoError(t, err)
		assert.Equal(t, expected, msg)
	})
}
