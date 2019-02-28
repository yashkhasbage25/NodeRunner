package server

import (
	"sync"

	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
)

var lock sync.Mutex

// BroadcastGameRedirection broadcasts the redirection msg to clients
// here, msg is to redirect the client to game page
func (server *Server) BroadcastGameRedirection() {
	firstWSocket := server.GetClient(0).GetWSocket()
	secondWSocket := server.GetClient(1).GetWSocket()

	redirector := dtypes.GameRedirector{Redirect: true}
	lock.Lock()
	firstWSocket.WriteJSON(redirector)
	lock.Unlock()
	lock.Lock()
	secondWSocket.WriteJSON(redirector)
	lock.Unlock()
	server.ClearConnections()
}

// ClearConnections closes the current websocket connections
func (server *Server) ClearConnections() {
	firstWSocket := server.GetClient(0).GetWSocket()
	secondWSocket := server.GetClient(1).GetWSocket()
	lock.Lock()
	firstWSocket.Close()
	lock.Unlock()
	lock.Lock()
	secondWSocket.Close()
	lock.Unlock()
	server.SetClient(0, nil)
	server.SetClient(1, nil)

	server.SetIDCounter(0)
}
