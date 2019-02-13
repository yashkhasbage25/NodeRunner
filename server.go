package server
 //checks weather  there are more than 1 client willing to join the game.

import(
	"client"
)

type server struct {
	ConnectedClient [2] *Client
	IdCounter int
}

func (s *server) validity() bool{
	if s.IdCounter != 2 {
		return false
	} else {
		return true
	}
}

func(s *server) AddNewClient (c *Client) {
	if s.IdCounter < 2 && c != nil {
		s.ConnectedClient[s.IdCounter] = c
		s.IdCounter++
	}
}

func (s *server) RemoveClient (IpAddress string) {
	if s.ConnectedClient[0].Ip == IpAddress {
		s.ConnectedClient[0] = s.ConnectedClient[1]
		s.ConnectedClient[1] = nil
		s.IdCounter--
	} else if s.ConnectedClient[1].Ip == IpAddress {
		s.ConnectedClient[1] = nil
		s.IdCounter--
	} else{
		//print wrong ip in consol.
	}
}

