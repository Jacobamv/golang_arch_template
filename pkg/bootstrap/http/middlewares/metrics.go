package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/bootstrap/http/misc/response/customRW"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/metrics"
)

func (m *provider) Metrics(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := customRW.NewResponseWriter(w)
		reqMethodAndPath := fmt.Sprintf("[%s] %s", r.Method, r.URL.Path)

		// HTTP Request Response Duration
		timer := prometheus.NewTimer(metrics.HttpRequestDuration.WithLabelValues(reqMethodAndPath))
		defer timer.ObserveDuration()

		next(lrw, r)

		// Determine microservice status based on response status code
		if lrw.Status() >= 200 && lrw.Status() < 500 {
			// Status 2xx or 3xx means microservice is UP
			metrics.SetMicroserviceStatus(1)
		} else {
			// Status 4xx or 5xx means microservice is DOWN
			metrics.SetMicroserviceStatus(0)
		}

		if r.URL.Path != "/api/metrics" {
			statusCode := strconv.Itoa(lrw.Status())

			// Same HTTP Request Path Counter
			metrics.HttpRequestCountWithPath.With(prometheus.Labels{"url": reqMethodAndPath}).Inc()
			metrics.HttpRequestStatusCode.With(prometheus.Labels{"code": statusCode, "url": reqMethodAndPath}).Inc()
			metrics.HttpRequestsPerTime.With(prometheus.Labels{"code": statusCode, "url": r.URL.Path}).Inc()
		}
	})
}
