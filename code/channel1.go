package main

import "fmt"

func main() {
	// Like maps and slices, channels must be created before use
	messages := make(chan string)

	// By default, sends and receives block until the other side is ready.
	go func() { messages <- "ping" }() // HL
	msg := <-messages                  // HL

	fmt.Println(msg)
}
