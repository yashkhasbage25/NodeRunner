package channels

import (
	dtypes "github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
)

// Data struct represents data that has to be communicated
// through channels while executing dijkstra
type Data struct {
	UpdatedPosition dtypes.Position
	MinimumDistance int
}

// Chans is an array of channels communicating through Data
var Chans [6]chan Data

// ChannelInitialization initializes channels
func ChannelInitialization() {
	for i := range Chans {
		Chans[i] = make(chan Data)
	}
}
