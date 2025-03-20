package router

import (
	"github.com/Jacobamv/golang_arch_template/internal/transport/amqp/consumers"
	"github.com/Jacobamv/golang_arch_template/pkg/brokers/rabbitmq"
)

func NewRouter(c *consumers.Consumers, client rabbitmq.Client) {
}
