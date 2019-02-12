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
