package main

import (
	dtypes "github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
	"log"
	"net/http"
	server "github.com/IITH-SBJoshi/concurrency-3/src/server"
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
