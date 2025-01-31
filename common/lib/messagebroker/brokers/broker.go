package brokers

import (
	"errors"

	"github.com/carlosealves2/go-file-converter/common/lib/messagebroker"
	"github.com/carlosealves2/go-file-converter/common/lib/messagebroker/brokers/rabbitmq"
)

type BrokerType string

var (
	RabbitMQBroker BrokerType = "rabbitmq"
)

func NewBroker(broker BrokerType, config any) (messagebroker.IMessageBroker, error) {
	switch broker {
	case RabbitMQBroker:
		rmqFactory := new(rabbitmq.RabbitMQFactory)
		return rmqFactory.CreateMessageBroker(config)
	default:
		return nil, errors.New("broker type not found")
	}
}
