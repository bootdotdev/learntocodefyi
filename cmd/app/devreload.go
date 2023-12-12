package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handlerDevReload(w http.ResponseWriter, r *http.Request) {
	log.Println("devreload handler connecting...")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if messageType == websocket.PingMessage {
			// Respond with a pong
			if err := conn.WriteMessage(websocket.PongMessage, nil); err != nil {
				break
			}
		}
	}
}
