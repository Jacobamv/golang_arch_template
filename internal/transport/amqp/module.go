package amqp

import (
	"gitlab.humo.tj/orzu/applications/bridge/internal/transport/amqp/consumers"
	"gitlab.humo.tj/orzu/applications/bridge/internal/transport/amqp/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	consumers.Module,
	router.Module,
)
