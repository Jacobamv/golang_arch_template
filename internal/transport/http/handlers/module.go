package handlers

import (
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewHandlerProvider)

// HandlerDependencies ...
type HandlerDependencies struct {
	fx.In
	Logger zerolog.Logger
}

// Handler ...
type Handler struct {
	logger zerolog.Logger
}

// NewHandlerProvider ...
func NewHandlerProvider(params HandlerDependencies) *Handler {
	return &Handler{
		logger: params.Logger,
	}
}
