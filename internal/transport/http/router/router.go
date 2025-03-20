package router

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.humo.tj/orzu/applications/bridge/internal/transport/http/handlers"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/bootstrap/http/middlewares"
	transportHTTP "gitlab.humo.tj/orzu/applications/bridge/pkg/bootstrap/http/router"
)

func AdaptHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

// NewRouter ..
func NewRouter(h *handlers.Handler, mw middlewares.Middleware) (router *transportHTTP.HTTPRouter) {
	router = transportHTTP.NewRouter()
	router.ConnectSwagger(h.ServeSwaggerFiles)

	router.GET("/ping", h.HPingPong, mw.RequestLog, mw.Metrics, mw.CORS)
	router.GET("/metrics", AdaptHandler(promhttp.Handler()), mw.RequestLog, mw.Metrics, mw.CORS)
	return
}
