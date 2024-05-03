package interceptor

import (
	"go.uber.org/zap"
	"homework/internal/events/service/producer"
	"homework/pkg/infrastructure/prometheus"
)

type Interceptor struct {
	logger   *zap.SugaredLogger
	producer *producer.EventsProducer
	metrics  *prometheus.ServerMetrics
}

func New(logger *zap.SugaredLogger, producer *producer.EventsProducer) *Interceptor {
	return &Interceptor{
		logger:   logger,
		producer: producer,
	}
}

func (i *Interceptor) SetMetrics(metrics *prometheus.ServerMetrics) {
	i.metrics = metrics
}
