package server

import (
	"dtypes"
)

// BroadcastGameRedirection broadcasts the redirection msg to clients
// here, msg is to redirect the client to game page
func (server *Server) BroadcastGameRedirection() {
	firstWSocket := server.GetClient(0).GetWSocket()
	secondWSocket := server.GetClient(1).GetWSocket()

	redirector := dtypes.GameRedirector{Redirect: true}

	firstWSocket.WriteJSON(redirector)
	secondWSocket.WriteJSON(redirector)

	server.ClearConnections()
}

// ClearConnections closes the current websocket connections
func (server *Server) ClearConnections() {
	firstWSocket := server.GetClient(0).GetWSocket()
	secondWSocket := server.GetClient(1).GetWSocket()

	firstWSocket.Close()
	secondWSocket.Close()

	server.SetClient(0, nil)
	server.SetClient(1, nil)

	server.SetIDCounter(0)
}
