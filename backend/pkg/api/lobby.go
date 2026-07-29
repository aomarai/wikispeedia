package api

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func CreateLobby(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a lobby
}

func GetLobby(w http.ResponseWriter, r *http.Request) {
	// Implementation for retrieving a lobby
}

var upgrader = websocket.Upgrader{} // use default options

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received: %s", message)

		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
