package main

import (
	"log"
	"time"

	nats "github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect("nats://nats-2:4222")
	if err != nil {
		log.Fatal(err)
	}
	msg, err := nc.Request("help", []byte("help me"), 100*time.Millisecond) // HL
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response received: %s", msg)
	nc.Close()
}
