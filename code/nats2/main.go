package main

import (
	"time"

	nats "github.com/nats-io/go-nats"
)

func main() {
	nc, _ := nats.Connect("nats://nats:4222") // HL
	// error handling omitted for simplicity
	for {
		nc.Publish("test", []byte("Hello Go Miami!")) // HL
		time.Sleep(time.Second * 1)
	}
}
