package client

import (
	"log"

	"github.com/gorilla/websocket"
)

// CompareClientsWithAddr compares the given ip and port with a client.
// if they have same ip and port then it returns true
func CompareClientsWithAddr(ip, port string, client *Client) bool {
	if client == nil {
		log.Fatal("CompareClientsWithAddr cannot compare when client is nil")
	}
	if client.GetIP() == ip && client.GetPort() == port {
		return true
	}
	return false
}

// CompareClientsWithSocket compares a given websocket connection with websocket
// connection of a given client
func CompareClientsWithSocket(wsocket *websocket.Conn, client *Client) bool {
	if wsocket == client.GetWSocket() {
		return true
	}
	return false
}
