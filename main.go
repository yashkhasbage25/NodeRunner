package main

import (
	"net/http"
	play "play_node_runner"
	server "server"
	"github.com/gorilla/websocket"
)

func main() {

	gameServer := server.Server{IdCounter: 0}

	gameServer.SetHandlers()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
