package client

import (
	dtypes "github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
	"strconv"

	websocket "github.com/gorilla/websocket"
)

// Client struct represents the connected client
type Client struct {
	IP             string
	ID             uint32
	Port           string
	WSocket        *websocket.Conn
	RequestChannel chan dtypes.Event
	ReceiveChannel chan dtypes.Event
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

// GetRequestChannel is getter for update channel of a client
func (client *Client) GetRequestChannel() chan dtypes.Event {
	return client.RequestChannel
}

// GetReceiveChannel is a getter for receive channel of client 
// object
func (client *Client) GetReceiveChannel() chan dtypes.Event {
	return client.ReceiveChannel
}

// GetInfoStr gives string of attributes of a client object
func (client *Client) GetInfoStr() string {
	return "client: IP: " + client.GetIP() + "\n\t\t" + "Port: " + client.GetPort() + "\n\t\t" + "ID: " + strconv.Itoa(int(client.GetID()))
}

// SetIP is getter for IP
func (client *Client) SetIP(ip string) {
	client.IP = ip
}

// SetID is getter for ID
func (client *Client) SetID(id uint32) {
	client.ID = id
}

// SetPort is getter for Port
func (client *Client) SetPort(port string) {
	client.Port = port
}

// SetWSocket is getter for Wsocket
func (client *Client) SetWSocket(wsocket *websocket.Conn) {
	client.WSocket = wsocket
}

// SetRequestChannel is setter for RequestChannel of client
func (client *Client) SetRequestChannel(requestChannel chan dtypes.Event) {
	client.RequestChannel = requestChannel
}

// SetReceiveChannel is a setter for receive channel of 
// a client object
func (client *Client) SetReceiveChannel(receiveChannel chan dtypes.Event) {
	client.ReceiveChannel = receiveChannel
}
