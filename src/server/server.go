package server

import "client"

type Server struct {
	ConnectedClient [2]*client.Client
	IdCounter uint32
}

func (s *Server) Validity() bool { 
	if s.IdCounter != 2 {
		return false
	} else {
		return true
	}
}

func(s *Server) AddNewClient (c *client.Client) {
	if s.IdCounter < 2 && c != nil {
		s.ConnectedClient[s.IdCounter] = c
		s.IdCounter++
	}
}

func (s *Server) RemoveClient (IpAddress string) {
	if s.ConnectedClient[0].Ip == IpAddress {
		s.ConnectedClient[0] = s.ConnectedClient[1]
		s.ConnectedClient[1] = nil
		s.IdCounter--
	} else if s.ConnectedClient[1].Ip == IpAddress {
		s.ConnectedClient[1] = nil
		s.IdCounter--
	} else {
		//print wrong ip in console.
	}
}
