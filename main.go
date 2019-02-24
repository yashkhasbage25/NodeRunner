package main

import (
	"fmt"
	"dijkstra"
	"dtypes"
	"coords"
)

func main() {
	bot:=dtypes.Position{15,25}
	player:=dtypes.Position{1185,530}
	nxtmove,distance:=dijkstra.RunDijkstra(bot,player,0)
	fmt.Println(nxtmove.X,"   ",nxtmove.Y)
	fmt.Println(coords.Platform)



}