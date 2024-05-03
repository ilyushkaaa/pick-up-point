package prometheus

import "github.com/prometheus/client_golang/prometheus"

type ServerMetrics struct {
	Hits                              prometheus.Counter
	HitsByResponseCodes               *prometheus.CounterVec
	HitsByResponseCodesAndRequestTime *prometheus.HistogramVec
}

func newServerMetrics() *ServerMetrics {
	return &ServerMetrics{
		Hits: createCounter("hits", "total num of hits to server"),
		HitsByResponseCodes: createCounterVec("hits_by_response_code",
			"total num of hits to server by response grpc code", "status_code", "request_method"),
		HitsByResponseCodesAndRequestTime: createHistogramVec("hits_by_response_codes_and_request_time",
			"total num of hits to server by response grpc code and response time",
			0, 0.005, 6, "status_code", "request_method"),
	}
}
