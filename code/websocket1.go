package main

import (
	"net/http"

	"github.com/gorilla/websocket" // HL
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // HL
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})
	http.ListenAndServe(":8080", nil)
}
