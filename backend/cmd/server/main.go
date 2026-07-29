package main

import (
	"log"
	"net/http"

	"wikispeedia/pkg/api"
)

func main() {
	http.HandleFunc("/api/v1/lobbies", api.CreateLobby)
	http.HandleFunc("/api/v1/lobbies/", api.GetLobby)
	http.HandleFunc("/api/v1/wiki/random", api.GetRandomArticles)
	http.HandleFunc("/ws/lobby/:id", api.WebSocketHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
