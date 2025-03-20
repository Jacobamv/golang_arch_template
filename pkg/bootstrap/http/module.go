package http

import (
	"gitlab.humo.tj/orzu/applications/bridge/pkg/bootstrap/http/middlewares"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/bootstrap/http/misc"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/bootstrap/http/server"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	middlewares.Module,
	misc.Module,

	server.ServerModule,
)
