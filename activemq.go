package goactivemq

import (
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type _ActiveMQ struct {
	showLog bool

	address  string
	port     uint16
	clientId string

	// Client from ActiveMQ (MQTT)
	client mqtt.Client
}

func (activemq *_ActiveMQ) SetLogger(showLog bool) {
	activemq.showLog = showLog
}

func (activemq *_ActiveMQ) SetClientId(clientId interface{}) {
	activemq.clientId = fmt.Sprint(clientId)
}

func (activemq *_ActiveMQ) Connect(address string, port uint16) (err error) {
	// Add options to connection
	opts := mqtt.NewClientOptions()

	activemq.address = address
	activemq.port = port

	activemq.setupOptions(opts)
	activemq.client = mqtt.NewClient(opts)
	if token := activemq.client.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
	}

	return err
}

func (activemq *_ActiveMQ) setupOptions(opts *mqtt.ClientOptions) {
	// Use tcp to comunication address
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", activemq.address, activemq.port))
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetCleanSession(false)

	if activemq.clientId != "" {
		opts.SetClientID(activemq.clientId)
	} else {
		opts.SetClientID(fmt.Sprint(rand.Intn(256)))
	}
}

func (activemq *_ActiveMQ) Disconnect() (err error) {
	activemq.client.Disconnect(250)
	return err
}

func (activemq *_ActiveMQ) Publisher(topic string, data []byte) (err error) {
	if token := activemq.client.Publish(topic, 0x01, false, data); token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	return err
}

func (activemq *_ActiveMQ) PublisherRetaining(topic string, data []byte) (err error) {
	if token := activemq.client.Publish(topic, 0x01, true, data); token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	return err
}

func (activemq *_ActiveMQ) Subscriber(topic string, handle func(data []byte)) (err error) {
	token := activemq.client.Subscribe(topic, 0x01, func(client mqtt.Client, msg mqtt.Message) {
		handle(msg.Payload())
	})

	token.WaitTimeout(16 * time.Second)

	return token.Error()
}
