package service

import (
	"gitlab.humo.tj/orzu/applications/bridge/internal/service/example"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	example.Module,
)
