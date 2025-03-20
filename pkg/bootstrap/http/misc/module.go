package misc

import (
	"gitlab.humo.tj/orzu/applications/bridge/pkg/bootstrap/http/misc/response"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	response.Module,
)
