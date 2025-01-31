package messagebroker

import "context"

type IMessageBroker interface {
	Publish(ctx context.Context, topic string, data []byte)
	Consume(ctx context.Context, consumer func(data []byte))
}
