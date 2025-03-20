package handlers

import (
	"github.com/Jacobamv/golang_arch_template/internal/service/example"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewHandlerProvider)

// HandlerDependencies ...
type HandlerDependencies struct {
	fx.In
	Logger  zerolog.Logger
	Example example.BExample
}

// Handler ...
type Handler struct {
	logger  zerolog.Logger
	example example.BExample
}

// NewHandlerProvider ...
func NewHandlerProvider(params HandlerDependencies) *Handler {
	return &Handler{
		logger:  params.Logger,
		example: params.Example,
	}
}
