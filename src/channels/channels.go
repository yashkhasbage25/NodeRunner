package channels

import(
	"dtypes"
	
)



type Data struct {
	UpdatedPosition dtypes.Position
	MinimumDistance int
}

var Chans [6]chan Data

func ChannelInitialization(){
	for i := range Chans {
   	Chans[i] = make(chan Data)
	}
}