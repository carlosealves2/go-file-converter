package rabbitmq

import (
	"runtime"

	"github.com/carlosealves2/go-file-converter/common/configs/brokersconfig"
	"github.com/carlosealves2/go-file-converter/common/lib/messagebroker"
	"github.com/phuslu/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQFactory struct{}

func (r *RabbitMQFactory) CreateMessageBroker(config any) (messagebroker.IMessageBroker, error) {
	rmqConfig := config.(brokersconfig.RabbitMQConfig)

	conn, err := amqp.Dial(rmqConfig.Uri)
	if err != nil {
		return handlerFactoryError(err, "failure while connect with RabbitMQ broker")
	}

	ch, err := conn.Channel()
	if err != nil {
		return handlerFactoryError(err, "failure while create channel with broker")
	}

	q, err := ch.QueueDeclare(
		rmqConfig.Queue.Name,
		rmqConfig.Queue.Durable,
		rmqConfig.Queue.AutoDelete,
		rmqConfig.Queue.Exclusive,
		rmqConfig.Queue.NoWait,
		rmqConfig.Queue.Arguments,
	)
	if err != nil {
		return handlerFactoryError(err, "failure to declare queue")
	}

	err = ch.ExchangeDeclare(
		rmqConfig.Exchange.Name,
		rmqConfig.Exchange.Type,
		rmqConfig.Exchange.Durable,
		rmqConfig.Exchange.AutoDelete,
		rmqConfig.Exchange.Internal,
		rmqConfig.Exchange.NoWait,
		rmqConfig.Exchange.Arguments,
	)
	if err != nil {
		return handlerFactoryError(err, "failure to declare exchange")
	}

	if rmqConfig.Key != "" {
		err = ch.QueueBind(q.Name, rmqConfig.Key, rmqConfig.Exchange.Name, false, nil)
		if err != nil {
			return handlerFactoryError(err, "failure to bind queue")
		}
	}

	return &RabbitMQBroker{
		conn:  conn,
		ch:    ch,
		queue: q.Name,
	}, nil
}

func handlerFactoryError(err error, msg string) (messagebroker.IMessageBroker, error) {
	_, file, line, _ := runtime.Caller(1)
	log.Error().Str("file", file).Int("line", line).Err(err).Msg(msg)
	return nil, err
}
