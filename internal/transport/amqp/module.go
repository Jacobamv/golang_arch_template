package amqp

import (
	"github.com/Jacobamv/golang_arch_template/internal/transport/amqp/consumers"
	"github.com/Jacobamv/golang_arch_template/internal/transport/amqp/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	consumers.Module,
	router.Module,
)
