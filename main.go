package main

import (
	"log"
	"net/http"
	server "server"
)

func main() {

	gameServer := server.Server{
		IDCounter: 0,
	}

	gameServer.SetHandlers(&gameServer)
	go gameServer.RedirectToGameIfConnected()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
