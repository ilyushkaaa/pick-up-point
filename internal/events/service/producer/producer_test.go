package producer

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework/tests/fixtures"
	"homework/tests/test_json"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()

	t.Run("unsuccessful send", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		s.mockProducer.EXPECT().SendMessage(gomock.Any()).Return(int32(1), int64(1), fmt.Errorf("error"))

		_, err := s.eventProducer.SendMessage(fixtures.Event().Valid().Value())

		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		expected := SendMessageResult{
			Partition: 1,
			Offset:    1,
		}
		s.mockProducer.EXPECT().SendMessage(gomock.Any()).Return(int32(1), int64(1), nil)

		msg, err := s.eventProducer.SendMessage(fixtures.Event().Valid().Value())

		assert.Equal(t, expected, msg)
		assert.NoError(t, err)
	})
}

func TestBuildMessage(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		s := setUp(t)
		expected := &sarama.ProducerMessage{
			Topic:     s.eventProducer.topic,
			Value:     sarama.ByteEncoder(test_json.ValidEventJSON),
			Partition: -1,
		}

		msg, err := s.eventProducer.BuildMessage(fixtures.Event().Valid().Value())

		assert.NoError(t, err)

		// можно было проверить просто assert.Equal(expected, msg)
		// но я посчитал что тогда тест упадет из за минимального различия в сериализации json поля Value
		// поэтому сделал более длинную проверку, но зато с assert.JSONEq
		valueJSONExpected, err := expected.Value.Encode()
		require.NoError(t, err)
		valueJSONActual, err := msg.Value.Encode()
		require.NoError(t, err)

		assert.JSONEq(t, string(valueJSONExpected), string(valueJSONActual))
		assert.Equal(t, expected.Topic, msg.Topic)
		assert.Equal(t, expected.Partition, msg.Partition)
	})
}
