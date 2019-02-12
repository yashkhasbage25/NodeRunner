package client

import (	
	"github.com/gorilla/websocket"
)
type Client struct {
	Ip string 
	Id string
	Wsocket *websocket.Conn
}