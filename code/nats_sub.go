package main

import (
	"fmt"
	"log"
	"time"

	nats "github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	nc.Subscribe("test", func(m *nats.Msg) { // HL
		fmt.Printf("Received a message: %s\n", string(m.Data)) // HL
	}) // HL
	time.Sleep(time.Second * 5)
}
