package main

import (
	"encoding/json"
	"flag"
	"time"

	"github.com/davidgaspardev/goactivemq"
)

var flagSession = flag.String("session", "my-id", "session of connectin with MQTT broker")
var flagAddress = flag.String("address", "10.1.15.59", "MQTT broker address")
var flagPort = flag.Int("port", 1883, "MQTT broker port")
var flagTopic = flag.String("topic", "PLC_COUNTER", "topic in the MQTT broker")
var flagQuantity = flag.Int("quantity", 500, "quantity of data to publish")

func main() {
	// Getting flags (arguments)
	session := *flagSession
	address := *flagAddress
	port := uint16(*flagPort)
	topic := *flagTopic
	quantity := *flagQuantity

	clientmq := goactivemq.NewActiveMQ()
	clientmq.SetClientId(session)

	if err := clientmq.Connect(address, port); err != nil {
		panic(err)
	}

	for i := 0; i < quantity; i++ {
		obj := map[string]interface{}{
			"resCode": "INJ003",
			"date":    time.Now().Format("2006-01-02 15:04:05"),
			"counter": i + 1,
		}

		data, err := json.Marshal(obj)
		if err != nil {
			panic(err)
		}

		if err := clientmq.Publisher(topic, data); err != nil {
			panic(err)
		}
	}

	clientmq.Disconnect()
}
