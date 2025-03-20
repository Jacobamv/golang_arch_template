package brokers

import (
	"context"

	"github.com/Jacobamv/golang_arch_template/pkg/brokers/rabbitmq"
	"github.com/Jacobamv/golang_arch_template/pkg/config"
	"github.com/rs/zerolog"
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
