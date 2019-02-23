package server

import (
	"client"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

// Server structure
type Server struct {
	ConnectedClient [2]*client.Client
	IDCounter       uint32
}

// GetIDCounter is getter for IDCounter of a server
func (server *Server) GetIDCounter() uint32 {
	return server.IDCounter
}

// GetClient is a getter for clients connected to server
func (server *Server) GetClient(index int) *client.Client {
	if index >= 2 {
		log.Println("Client index requested is out of bounds")
		return nil
	}
	return server.ConnectedClient[index]
}

// SetIDCounter is setter for IDCoubter
func (server *Server) SetIDCounter(i uint32) {
	server.IDCounter = i
}

// SetClient is a setter for client of server
func (server *Server) SetClient(index uint32, newClient *client.Client) {
	server.ConnectedClient[index] = newClient
}

// CheckClientLimit checks if the number of clients is not more than 2
func (server *Server) CheckClientLimit() bool {
	if server.GetIDCounter() <= 2 {
		return true
	} else {
		return false
	}
}

// IncrementIDCounter increments the IDCounter by 1
func (server *Server) IncrementIDCounter() {
	server.IDCounter++
}

// AddNewClient adds a new client to the gameServer server
func (server *Server) AddNewClient(newClient *client.Client) {
	if newClient == nil {
		log.Fatal("Cannot add client because Client is nil")
	}
	if server.GetIDCounter() < 2 {
		server.SetClient(server.GetNextID(), newClient)
		server.IncrementIDCounter()
	}
	log.Println("Added new client to server", newClient.GetInfoStr())
}

// RemoveClientWithAddr removes the client given ots IPAddress,
// in case of some disconnections.
func (server *Server) RemoveClientWithAddr(ipaddress, port string) {

	firstClient := server.ConnectedClient[0]
	secondClient := server.ConnectedClient[1]

	if client.CompareClientsWithAddr(ipaddress, port, firstClient) {
		server.SetClient(1, nil)
		server.SetClient(0, secondClient)
	} else if client.CompareClientsWithAddr(ipaddress, port, secondClient) {
		server.SetClient(1, nil)
	} else {
		log.Printf("Server does not contain client with ip:" + ipaddress + " port:" + port)
	}
}

// RemoveClientWithSocket removes the client given its websocket connection
func (server *Server) RemoveClientWithSocket(socket *websocket.Conn) {
	firstClient := server.GetClient(0)
	secondClient := server.GetClient(1)

	if client.CompareClientsWithSocket(socket, firstClient) {
		server.SetClient(1, nil)
		server.SetClient(0, secondClient)
	} else if client.CompareClientsWithSocket(socket, secondClient) {
		server.SetClient(1, nil)
	} else {
		log.Printf("Server does not contain client with given websocket:")
	}
}

// GetNextID returns the next possible ID that can be given to a client
func (server *Server) GetNextID() uint32 {
	return server.IDCounter
}

// AreBothClientsConnected checks if both clients are connected or not
func (server *Server) AreBothClientsConnected() bool {
	firstClient := server.GetClient(0)
	secondClient := server.GetClient(1)

	if firstClient != nil && secondClient != nil {
		return true
	}
	return false
}

// RedirectToGameIfConnected redirects both clients to game.html once both clients
// are connected
func (server *Server) RedirectToGameIfConnected() {
	log.Println("begin execution of RedirectToGameIfConnected")
	for {
		if server.AreBothClientsConnected() {
			break
		}
	}
	log.Println("Both clients connected to server")
	server.BroadcastGameRedirection()
}

// SetConnWithAddr sets the websocket connection of a cleint object given its ip and port
func (server *Server) SetConnWithAddr(ip, port string, conn *websocket.Conn) uint32 {
	if client.CompareClientsWithAddr(ip, port, server.GetClient(0)) {
		server.GetClient(0).SetWSocket(conn)
		return 0
	} else if client.CompareClientsWithAddr(ip, port, server.GetClient(1)) {
		server.GetClient(1).SetWSocket(conn)
		return 1
	} else {
		log.Fatalf("No client found with ip: " + ip + " port: " + port)
		return 2
	}
}

// GetInfoStr gives a string  of attributes of a server object
func (server *Server) GetInfoStr() string {
	firstClient := server.GetClient(0)
	secondClient := server.GetClient(1)
	clientInfo := "clients: "
	if firstClient != nil {
		clientInfo += " client1: " + firstClient.GetInfoStr()
	}
	if secondClient != nil {
		clientInfo += " client2: " + secondClient.GetInfoStr()
	}
	return "server: " + "IDCounter: " + strconv.Itoa(int(server.GetIDCounter())) + " " + clientInfo
}
