package response

import (
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewResponse)

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger zerolog.Logger
}

// Response ...
type Response struct {
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`

	contentType string `json:"-"`
	logger      zerolog.Logger
}

// NewResponse ...
func NewResponse(params Dependencies) *Response {
	return &Response{
		logger: params.Logger,
	}
}
