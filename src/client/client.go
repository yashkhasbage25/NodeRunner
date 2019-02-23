package client

import (
	"strconv"

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

// GetWSocket is getter for Wsocket
func (client *Client) GetWSocket() *websocket.Conn {
	return client.WSocket
}

// GetInfoStr gives string of attributes of a client object
func (client *Client) GetInfoStr() string {
	return "client: IP: " + client.GetIP() + "\n\t\t" + "Port: " + client.GetPort() + "\n\t\t" + "ID: " + strconv.Itoa(int(client.GetID()))
}
