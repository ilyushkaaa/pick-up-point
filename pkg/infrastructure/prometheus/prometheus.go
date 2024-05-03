package prometheus

import (
	"net/http"
	"os"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Init(s *grpc.Server, logger *zap.SugaredLogger, grpcMetrics *grpc_prometheus.ServerMetrics) (*BusinessMetrics, *ServerMetrics) {
	reg := prometheus.NewRegistry()
	reg.MustRegister(grpcMetrics)
	grpcMetrics.InitializeMetrics(s)

	bm := newBusinessMetrics()
	sm := newServerMetrics()

	reg.MustRegister(bm.OrdersCountByPickUpPoints, bm.DeliveredOrdersCount, bm.IssuedOrdersCount,
		bm.ReturnedOrdersCount, bm.OrdersByClientCount, bm.DeletedOOrders, sm.Hits, sm.HitsByResponseCodes,
		sm.HitsByResponseCodesAndRequestTime)

	go func() {
		logger.Fatal(http.ListenAndServe(os.Getenv("METRICS_ADDR"), promhttp.HandlerFor(reg, promhttp.HandlerOpts{EnableOpenMetrics: true})))
	}()

	return bm, sm
}

func createCounterVec(name, help string, labels ...string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: name,
		Help: help,
	}, labels)
}

func createHistogramVec(name, help string, minVal, interval float64, bucketsNum int, labels ...string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    name,
		Help:    help,
		Buckets: prometheus.LinearBuckets(minVal, interval, bucketsNum),
	}, labels)
}

func createCounter(name, help string) prometheus.Counter {
	return prometheus.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: help,
	})
}
