package server
 //checks weather  there are more than 1 client willing to join the game.

import(
	"client"
)

type server struct {
		Connected_client [2] *Client
		Id_counter int

}

func (s server) validity() bool{
	if s.Id_counter !=2	{
		return false
	}else{
		return true
	}
}

func(s server) AddNewClient(c *Client)
{
	if s.Id_counter<2 && c!=nil{
		s.Connected_client[s.Id_counter]=c
		s.Id_counter++
	}

}
func (s server) RemoveClient(IpAddress string)
{
	if s.Connected_client[0].Ip==IpAddress{
		s.Connected_client[0]=s.Connected_client[1]
		s.Connected_client[1]=nil
		s.Id_counter--
	}
	else if s.Connected_client[1].Ip==IpAddress{
		s.Connected_client[1]=nil
		s.Id_counter--
	}
	else{
		//print wrong ip in consol.
	}

}

