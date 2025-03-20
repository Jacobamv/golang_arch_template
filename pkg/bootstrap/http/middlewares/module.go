package middlewares

import (
	"net/http"

	"github.com/Jacobamv/golang_arch_template/pkg/config"
	"github.com/rs/zerolog"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewMiddleware)

// Middleware ...
type Middleware interface {
	// Сбор метрик сервиса
	Metrics(next http.HandlerFunc) http.HandlerFunc
	CORS(next http.HandlerFunc) http.HandlerFunc
	RequestLog(next http.HandlerFunc) http.HandlerFunc
}

// Dependencies ...
type Dependencies struct {
	fx.In

	Config *config.Config
	Logger zerolog.Logger
}

type provider struct {
	config *config.Config
	logger zerolog.Logger
}

// NewMiddleware ...
func NewMiddleware(params Dependencies) Middleware {
	return &provider{
		config: params.Config,
		logger: params.Logger,
	}
}
