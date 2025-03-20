package brokers

import (
	"context"

	"github.com/rs/zerolog"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/brokers/rabbitmq"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/config"
	"go.uber.org/fx"
)

// AMQPModule ...
var Module = fx.Provide(NewAMQPConn)

// Dependencies ...
type Dependencies struct {
	fx.In

	Logger zerolog.Logger
	Config *config.Config
}

// NewAMQPConn ...
func NewAMQPConn(params Dependencies) rabbitmq.Client {
	return NewRabbitMQConn(params, context.Background())
}
