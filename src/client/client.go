package client

import (
	"github.com/gorilla/websocket"
)

// Client struct represents the connected client
type Client struct {
	IP      string
	ID      uint32
	Port    string
	WSocket *websocket.Conn
}

// GetIP is getter for IP
func (client *Client) GetIP() string {
	return client.IP
}

// GetID is getter for ID
func (client *Client) GetID() uint32 {
	return client.ID
}

// GetPort is getter for Port
func (client *Client) GetPort() string {
	return client.Port
}

// GetWsocket is getter for Wsocket
func (client *Client) GetWSocket() *websocket.Conn {
	return client.WSocket
}
