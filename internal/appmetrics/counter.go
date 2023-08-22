package appmetrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	GRPCPanicsCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "device_data_api_panics_total",
		Help: "Total Panics recovered",
	})

	GRPCRequestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "device_data_api_grpc_request_count",
			Help: "The total number of requests served by the GRPC Server",
		},
		[]string{"method", "status"},
	)

	GRPCResponseTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "device_data_api_grpc_response_time",
			Help:    "The response time distribution of the GRPC Server",
			Buckets: []float64{0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		},
		[]string{"method", "status"},
	)

	HTTPRequestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "device_data_api_http_request_count",
			Help: "The total number of requests served by the Http Server",
		},
		[]string{"method", "path", "status"},
	)

	HTTPResponseTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "device_data_api_http_response_time",
			Help:    "The response time distribution of the Http Server",
			Buckets: []float64{0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		},
		[]string{"method", "path", "status"},
	)
)
