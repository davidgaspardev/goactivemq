package main

import (
	"flag"
	"fmt"

	"github.com/davidgaspardev/goactivemq"
)

var flagSession = flag.String("session", "my-id", "session of connectin with MQTT broker")
var flagAddress = flag.String("address", "10.1.15.59", "MQTT broker address")
var flagPort = flag.Int("port", 1883, "MQTT broker port")
var flagTopic = flag.String("topic", "PLC_COUNTER", "topic in the MQTT broker")

func main() {
	// Getting flags (arguments)
	session := *flagSession
	address := *flagAddress
	port := uint16(*flagPort)
	topic := *flagTopic

	clientMQ := goactivemq.NewActiveMQ()
	clientMQ.SetClientId(session)

	fmt.Println("Starting connection")
	if err := clientMQ.Connect(address, port); err != nil {
		panic(err)
	}

	receptor := make(chan string)

	fmt.Println("Send subscribe")
	if err := clientMQ.Subscriber(topic, func(data []byte) {
		receptor <- string(data)
	}); err != nil {
		panic(err)
	}

	fmt.Println("Listening", topic)

	for {
		data := <-receptor
		fmt.Println(data)
	}
}
