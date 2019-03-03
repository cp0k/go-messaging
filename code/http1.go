package main

import (
	"fmt"
	"log"
	"net/http" // HL
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)                // HL
	log.Fatal(http.ListenAndServe(":8080", nil)) // HL
}
