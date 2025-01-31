package messagebroker

type IMessageBrockerFactory interface {
	CreateMessageBroker(config any) (IMessageBroker, error)
}
