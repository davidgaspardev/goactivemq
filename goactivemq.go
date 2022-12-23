package goactivemq

type ActiveMQ interface {
	// Debug
	SetLogger(showLog bool)

	SetClientId(session interface{})

	Connect(address string, port uint16) error
	Disconnect() error

	Publisher(topic string, data []byte) error
	PublisherRetaining(topic string, data []byte) error

	Subscriber(topic string, handle func(data []byte)) error
}

func NewActiveMQ() ActiveMQ {
	return &_ActiveMQ{}
}
