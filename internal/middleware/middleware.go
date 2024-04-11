package middleware

import (
	"go.uber.org/zap"
	"homework/internal/events/service/producer"
)

type Middleware struct {
	logger   *zap.SugaredLogger
	producer *producer.EventsProducer
}

func New(logger *zap.SugaredLogger, producer *producer.EventsProducer) *Middleware {
	return &Middleware{
		logger:   logger,
		producer: producer,
	}
}
