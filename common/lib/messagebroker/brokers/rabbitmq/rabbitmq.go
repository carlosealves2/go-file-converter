package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQBroker struct {
	conn            *amqp.Connection
	ch              *amqp.Channel
	exchange, queue string
}

func (r *RabbitMQBroker) Publish(ctx context.Context, topic string, data []byte) {}

func (r *RabbitMQBroker) Consume(ctx context.Context, consumer func(data []byte)) {}
