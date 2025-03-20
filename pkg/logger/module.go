package logger

import (
	"gitlab.humo.tj/orzu/applications/bridge/pkg/config"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(InitLogger)

// Dependencies ...
type Dependencies struct {
	fx.In

	Config *config.Config
}
