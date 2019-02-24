package main

import (
	"dtypes"
	"log"
	"net/http"
	server "server"
)

func main() {

	gameServer := server.Server{
		IDCounter:      0,
		RequestChannel: make(chan dtypes.Events),
	}

	gameServer.SetHandlers()
	go gameServer.RedirectToGameIfConnected()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
