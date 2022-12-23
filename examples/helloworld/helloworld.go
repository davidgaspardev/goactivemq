package main

import (
	"encoding/json"
	"flag"

	"github.com/davidgaspardev/goactivemq"
)

var flagAddress = flag.String("address", "10.1.15.59", "MQTT broker address")
var flagPort = flag.Int("port", 1883, "MQTT broker port")
var flagTopic = flag.String("topic", "PLC_COUNTER", "topic in the MQTT broker")

func main() {
	// Getting flags (arguments)
	address := *flagAddress
	port := uint16(*flagPort)
	topic := *flagTopic

	clientMQ := goactivemq.NewActiveMQ()
	if err := clientMQ.Connect(address, port); err != nil {
		panic(err)
	}

	obj := map[string]interface{}{
		"hello": 0x776f726c64,
	}

	data, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	if err := clientMQ.Publisher(topic, data); err != nil {
		panic(err)
	}

	clientMQ.Disconnect()
}
