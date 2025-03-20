package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Defining constants for labels
	LabelURL  = "url"
	LabelPath = "path"
	LabelCode = "code"

	HttpRequestCountWithPath = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "example_http_requests_total_with_path",
			Help: "Number of HTTP requests by path.",
		},
		[]string{LabelURL},
	)

	// PROMQL => rate(http_request_duration_seconds_sum{}[5m]) / rate(http_request_duration_seconds_count{}[5m])
	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "example_http_request_duration_seconds",
			Help: "Response time of HTTP request.",
		},
		[]string{LabelPath},
	)

	HttpRequestStatusCode = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "example_http_request_status_code_with_path",
			Help: "Count status code of HTTP request.",
		},
		[]string{LabelCode, LabelURL},
	)

	HttpRequestsPerTime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "example_http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{LabelCode, LabelURL},
	)

	MicroserviceStatus = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "example_status",
			Help: "Status of the microservice (1 = UP, 0 = DOWN).",
		},
	)
)

// SetMicroserviceStatus sets the status of the microservice
func SetMicroserviceStatus(status float64) {
	MicroserviceStatus.Set(status)
}

func init() {
	// Grouping metric registration
	prometheus.MustRegister(
		HttpRequestCountWithPath,
		HttpRequestDuration,
		HttpRequestStatusCode,
		HttpRequestsPerTime,
		MicroserviceStatus,
	)
}
