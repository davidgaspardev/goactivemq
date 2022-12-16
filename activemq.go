package activemq

import (
	"github.com/go-stomp/stomp/v3"
	"github.com/go-stomp/stomp/v3/frame"
)

type ActiveMQ interface {
	// Debug
	SetLogger(showLog bool)

	Connect(address string) error
	Disconnect() error

	Publisher(topic string, msg string) error
	Subscriber(topic string, handle func(msg string, err error)) error
}

type _ActiveMQ struct {
	showLog bool
	addr    string

	conn *stomp.Conn
}

func NewActiveMQ() ActiveMQ {
	return &_ActiveMQ{}
}

func (activemq *_ActiveMQ) SetLogger(showLog bool) {
	activemq.showLog = showLog
}

func (activemq *_ActiveMQ) Connect(address string) (err error) {
	// Add options to connection
	var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
		stomp.ConnOpt.Login("admin", "admin"),
	}

	activemq.conn, err = stomp.Dial("tcp", address, options...)
	return err
}

func (activemq *_ActiveMQ) Disconnect() (err error) {
	if activemq.conn != nil {
		err = activemq.Disconnect()
	}
	return err
}

func (activemq *_ActiveMQ) Publisher(topic string, msg string) (err error) {
	err = activemq.conn.Send(topic, "text/plain", []byte(msg), func(framing *frame.Frame) error {
		// Add RECEIPT header to sync the communication
		framing.Header.Add(frame.Receipt, frame.Id)
		return nil
	})
	return err
}

func (activemq *_ActiveMQ) Subscriber(topic string, handle func(msg string, err error)) error {
	subscribe, err := activemq.conn.Subscribe(topic, stomp.AckAuto)
	if err != nil {
		return err
	}

	defer subscribe.Unsubscribe()

	for {
		msg := <-subscribe.C
		err = msg.Err
		handle(string(msg.Body), err)
	}
}
