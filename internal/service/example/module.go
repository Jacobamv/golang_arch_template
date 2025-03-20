package example

import (
	"context"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBExample)

type BExample interface {
	Example(ctx context.Context) (err error)
}

type provider struct {
	// Logger

	logger zerolog.Logger
}

type Params struct {
	fx.In

	// Logger
	Logger zerolog.Logger
}

// NewBExample ...
func NewBExample(params Params) BExample {
	return &provider{
		// Logger
		logger: params.Logger,
	}
}
