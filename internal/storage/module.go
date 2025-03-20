package storage

import (
	"gitlab.humo.tj/orzu/applications/bridge/internal/storage/example"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	example.Module,
)
