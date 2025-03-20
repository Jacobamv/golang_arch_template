package gateways

import (
	"gitlab.humo.tj/orzu/applications/bridge/pkg/gateways/example"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	example.Module,
)
