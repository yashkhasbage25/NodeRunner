package main

import (
	"net/http"
	"log"
	server "server"
	
)

func main() {

	gameServer := server.Server{IdCounter: 0}

	gameServer.SetHandlers()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
