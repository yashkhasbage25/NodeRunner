package client

import (
	"testing"

	"github.com/gorilla/websocket"
)

func TestCompareClientsWithAddr(t *testing.T) {
	samples := []struct {
		x    string
		port string
		c    *Client
	}{
		{x: "192.0.0.1", port: "9000", c: &Client{IP: "192.1.0.0", Port: "9000"}},
		{x: "127.0.0.1", port: "80", c: &Client{IP: "192.168.0.1", Port: "80"}},
		{x: "192.168.0.1", port: "9000", c: &Client{IP: "192.168.0.1", Port: "9000"}},
		{x: "192.168.0.01", port: "10000", c: &Client{IP: "192.168.0.1", Port: "10000"}},
	}
	answers := []bool{false, false, true, false}
	for i := 0; i < len(answers); i++ {
		got := CompareClientsWithAddr(samples[i].x, samples[i].port, samples[i].c)
		want := answers[i]
		if got != want {
			t.Error("CompareClientsWithLadder was incorrect, got: ", got, ", want: ", want, " for x:", samples[i].x, " port: ", samples[i].port, " c: ", samples[i].c.GetInfoStr())
		}
	}
}

func TestCompareClientsWithSocket(t *testing.T) {
	sockets := []*websocket.Conn{
		&websocket.Conn{},
		&websocket.Conn{},
	}
	samples := []struct {
		socket *websocket.Conn
		c      *Client
	}{
		{socket: sockets[0], c: &Client{WSocket: sockets[1]}},
		{socket: sockets[1], c: &Client{WSocket: sockets[1]}},
		{socket: sockets[1], c: &Client{WSocket: sockets[0]}},
		{socket: sockets[0], c: &Client{WSocket: sockets[0]}},
	}

	answers := []bool{false, true, false, true}
	for i := 0; i < len(answers); i++ {
		got := CompareClientsWithSocket(samples[i].socket, samples[i].c)
		want := answers[i]
		if got != want {
			t.Error("CompareClientsWithSocket was incorrect, got:", got, "want:", want, "for ", samples[i])
		}
	}
}
