package main

import (
	"log"
	"os"
	"os/signal"

	nats "github.com/nats-io/go-nats"
)

func serviceWorker() *nats.Conn {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	nc.Subscribe("help", func(m *nats.Msg) { // HL
		nc.Publish(m.Reply, []byte("I can help!")) // HL
	}) // HL
	return nc
}

func main() {
	nc := serviceWorker() // HL
	defer nc.Close()      // HL

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}
