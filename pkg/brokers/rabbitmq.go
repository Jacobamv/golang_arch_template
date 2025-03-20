package brokers

import (
	"context"
	"fmt"

	"github.com/Jacobamv/golang_arch_template/pkg/brokers/rabbitmq"
)

const delay = 3 // reconnect after delay seconds

// NewNewRabbitMQ ...
func NewRabbitMQConn(params Dependencies, ctx context.Context) rabbitmq.Client {
	rmqp := fmt.Sprintf("amqp://%v:%v@%v:%v/",
		params.Config.RabbitMQ.Username,
		params.Config.RabbitMQ.Password,
		params.Config.RabbitMQ.Host,
		params.Config.RabbitMQ.Port,
	)

	client, err := rabbitmq.NewClient(rmqp)

	if err != nil {
		params.Logger.Fatal().Err(err).Send()
		return nil
	}

	return client

}
