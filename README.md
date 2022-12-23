# Go ActiveMQ

Go ActiveMQ is a cliente MQTT.

## Installation

Add package to your Go project:

```bash
$ go get github.com/davidgaspardev/goactivemq
```

## Examples

Send hello world object to HELLO_WORLD topic:

```go
package main

import (
	"encoding/json"

	"github.com/davidgaspardev/goactivemq"
)

func main() {
	// Getting flags (arguments)
	address := "10.1.15.59"
	port := uint16(1883)
	topic := "HELLO_WORLD"

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
```