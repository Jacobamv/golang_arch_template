package router

import (
	"gitlab.humo.tj/orzu/applications/bridge/internal/transport/amqp/consumers"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/brokers/rabbitmq"
)

func NewRouter(c *consumers.Consumers, client rabbitmq.Client) {
}
