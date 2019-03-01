package main

import (
	"io/ioutil"
	"log"
	"net/http"

	dtypes "github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
	server "github.com/IITH-SBJoshi/concurrency-3/src/server"
)

func main() {
	log.SetOutput(ioutil.Discard)
	gameServer := server.Server{
		IDCounter:      0,
		RequestChannel: make(chan dtypes.Event),
	}

	gameServer.SetHandlers()
	go gameServer.RedirectToGameIfConnected()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
