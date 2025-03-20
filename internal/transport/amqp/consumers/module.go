package consumers

import (
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewConsumers)

type Consumers struct {
	logger zerolog.Logger
}

type Dependencies struct {
	fx.In
	Logger zerolog.Logger
}

func NewConsumers(params Dependencies) *Consumers {
	return &Consumers{
		logger: params.Logger,
	}
}
