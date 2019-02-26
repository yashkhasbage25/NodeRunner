package dijkstra

import (
	"log"

	"github.com/IITH-SBJoshi/concurrency-3/src/channels"
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
	"github.com/IITH-SBJoshi/concurrency-3/src/platform"
)

// StaticNode number of static nodes are 32 depended on construction of graph ..here manually measured
type StaticNode struct {
	Location dtypes.Position
	NodeID   int
	XNodeID  int
	YNodeID  int
	//type of node ..ladder or ..end of path or ..
}

type Matrix struct { // 32 is number of static node and we are adding 2 dynamic node.
	AdjacencyMatrix [6][32 + 2][32 + 2]int
	Size            int // 34

}

// global game Matrix
var Game Matrix
var Parentarray [6][32 + 2]int
var Allnodes [6][32 + 2]StaticNode
var Path [6][32 + 2]bool

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func setter(i int, n int, z int, flag bool) {
	if flag == true {
		Game.AdjacencyMatrix[z][n][i] = Abs(Allnodes[z][i].Location.Y - Allnodes[z][n].Location.Y)
		Game.AdjacencyMatrix[z][i][n] = Abs(Allnodes[z][i].Location.Y - Allnodes[z][n].Location.Y)
	} else {
		Game.AdjacencyMatrix[z][n][i] = Abs(Allnodes[z][i].Location.X - Allnodes[z][n].Location.X)
		Game.AdjacencyMatrix[z][i][n] = Abs(Allnodes[z][i].Location.X - Allnodes[z][n].Location.X)
	}

}

func GetBoundary(player dtypes.Position) dtypes.Rect {
	var temp dtypes.Rect
	temp.XHi = player.X - 20
	temp.YHi = player.Y - 15
	temp.XLo = player.X + 20
	temp.YLo = player.Y + 15
	return temp
}
func AllignedWithLadder(player dtypes.Rect) int {
	var i int
	var center int = (player.XLo + player.XHi) / 2
	for i = 0; i < len(platform.Ladder); i++ {
		if platform.Ladder[i].YLo >= player.YLo && platform.Ladder[i].YHi <= player.YLo {
			if center >= platform.Ladder[i].XHi && center <= platform.Ladder[i].XLo {
				log.Println("AllignedWithLadder returns true.")
				return 1
			}
		}
	}
	log.Println("AllignedWithLadder returns false.")
	return 0
}

func OnPlatform(player dtypes.Rect) int {
	var i int
	log.Println("Executing OnPlatform")
	for i = 0; i < len(platform.Platform); i++ {
		//log.Println(player.XHi,platform.Platform[i].XHi, "---", player.YLo,"---",platform.Platform[i].YHi, "---", player.XLo,platform.Platform[i].XLo)
		if player.YLo == platform.Platform[i].YHi && player.XLo > platform.Platform[i].XHi && player.XHi < platform.Platform[i].XLo {
			log.Println("OnPlatform returns true.")
			return 1
		}
	}
	log.Println("OnPlatform returns false.")
	return 0
}
func fallingon(entity dtypes.Position, z int) int {
	var ymin = 1200
	for i := 0; i < 32; i++ {
		temp := dtypes.Position{entity.X, Allnodes[z][i].Location.Y}
		if Allnodes[z][i].Location.Y > entity.Y && Allnodes[z][i].Location.Y < ymin && OnPlatform(GetBoundary(temp)) == 1 {
			ymin = Allnodes[z][i].Location.Y
		}

	}
	return ymin
}
func onladder(entity dtypes.Position) bool { //code from atharva.
	output := AllignedWithLadder(GetBoundary(entity))
	if output == 1 {
		return true
	} else {
		return false
	}
}

func addDynamicnode(bot dtypes.Position, player dtypes.Position, z int) {

	Allnodes[z][32] = StaticNode{dtypes.Position{bot.X, bot.Y}, 9, bot.X, bot.Y} //added bot at position equal to n.
	if OnPlatform(GetBoundary(player)) == 0 && AllignedWithLadder(GetBoundary(player)) == 0 {
		Allnodes[z][32+1] = StaticNode{dtypes.Position{player.X, fallingon(player, z)}, 10, player.X, fallingon(player, z)}
	} else {
		Allnodes[z][32+1] = StaticNode{dtypes.Position{player.X, player.Y}, 10, player.X, player.Y} //added player
	}
	botonladder := AllignedWithLadder(GetBoundary(bot))
	playeronladder := AllignedWithLadder(GetBoundary(player))
	for i := 0; i < 32; i++ {
		if Allnodes[z][i].XNodeID == Allnodes[z][32].XNodeID && botonladder == 1 {
			setter(i, 32, z, true)
		} else if Allnodes[z][i].YNodeID == Allnodes[z][32].YNodeID {
			setter(i, 32, z, false)
		}
	}
	for i := 0; i < 32+1; i++ {
		if Allnodes[z][i].XNodeID == Allnodes[z][32+1].XNodeID && playeronladder == 1 {
			setter(i, 32+1, z, true)
		} else if Allnodes[z][i].YNodeID == Allnodes[z][32+1].YNodeID { // if it is not on ladder do not add edge if y> bot
			setter(i, 32+1, z, false) //position
		}
	}

}
func minDistance(distance []int, cluster []bool, size int) int {

	var min_index int
	min := int(^uint(0) >> 1)
	for v := 0; v < size; v++ { // we have to impliment this parallaly
		if cluster[v] == false && distance[v] <= min {
			min = distance[v]
			min_index = v
		}
	}
	// log.Println("some min distacne is", min_index)
	return min_index
}
func printPath(z int, node int) {

	if Parentarray[z][node] != -1 {
		if Parentarray[z][node] != 32 {
			printPath(z, Parentarray[z][node])
		}
		log.Println(node)
	}
}
func markPath(z int, node int) {

	if Parentarray[z][node] != -1 {
		if Parentarray[z][node] != 32 {
			markPath(z, Parentarray[z][node])
		}
		Path[z][node] = true
	}
}
func nextposition(currentPosition dtypes.Position, botNextmove dtypes.Position, step int) dtypes.Position {
	var updatedPosition dtypes.Position
	xcurrent := currentPosition.X
	ycurrent := currentPosition.Y
	xpropoposed := botNextmove.X
	ypropoposed := botNextmove.Y
	if xcurrent == xpropoposed {
		updatedPosition.X = xcurrent
		if ypropoposed > ycurrent {
			updatedPosition.Y = ycurrent + step
		} else if ypropoposed < ycurrent {
			updatedPosition.Y = ycurrent - step
		} else {
			updatedPosition.Y = ycurrent
		}
	} else if ycurrent == ypropoposed {
		updatedPosition.Y = ycurrent
		if xpropoposed > xcurrent {
			updatedPosition.X = xcurrent + step
		} else {
			updatedPosition.X = xcurrent - step
		}
	}
	return updatedPosition
}
func minimum(distance []int, i int, j int) int {
	if distance[i] < distance[j] {
		return i
	} else {
		return j
	}
}
func UpdateBots(event dtypes.Event) dtypes.Event {
	replyEvent := event
	minpathlen := make([]int, 6)
	var update [6]dtypes.Position
	var bestUpdate [3]dtypes.Position
	log.Println("Inside Updatebots eith event", event.GetStr())
	go runDijkstra(event.B1Pos, event.P1Pos, 0, channels.Chans[0])
	go runDijkstra(event.B1Pos, event.P2Pos, 1, channels.Chans[1])
	go runDijkstra(event.B2Pos, event.P1Pos, 2, channels.Chans[2])
	go runDijkstra(event.B2Pos, event.P2Pos, 3, channels.Chans[3])
	go runDijkstra(event.B3Pos, event.P1Pos, 4, channels.Chans[4])
	go runDijkstra(event.B3Pos, event.P2Pos, 5, channels.Chans[5])
	log.Println("Total 6 dijkstra are initiated")
	for i := 0; i < 6; i++ {
		var channeldata channels.Data
		channeldata = (<-channels.Chans[i])
		update[i] = channeldata.UpdatedPosition
		minpathlen[i] = channeldata.MinimumDistance
	}
	log.Println("In updated bots, obtained all channels after completing dijkstra")
	for i := 0; i < 3; i++ {
		bestUpdate[i] = update[minimum(minpathlen, 2*i, 2*i+1)]
	}
	log.Println("In update bots best updates are", bestUpdate[0].GetStr(), bestUpdate[1].GetStr(), bestUpdate[2].GetStr())
	replyEvent.B1Pos = bestUpdate[0]
	replyEvent.B2Pos = bestUpdate[1]
	replyEvent.B3Pos = bestUpdate[2]
	log.Println("Update bots replies with event", replyEvent.B1Pos.GetStr(), replyEvent.B2Pos.GetStr(), replyEvent.B3Pos.GetStr())
	return replyEvent

}
func runDijkstra(bot dtypes.Position, player dtypes.Position, z int, channel chan channels.Data) {
	var step int
	step = 2
	//initialize()
	addDynamicnode(bot, player, z)
	//we have source as node with NodeID
	//var distance [32+2] int
	distance := make([]int, 34)
	cluster := make([]bool, 34)
	//var cluster [32+2] bool // cluster[i] will be true if node i is included in shortest
	// path tree or shortest distance from src to i is finalize

	for i := 0; i < 32+2; i++ {
		distance[i] = int(^uint(0) >> 1) // it is int_max in golang
		cluster[i] = false
	}
	distance[32] = 0 // distance of src from src is 0

	for i := 0; i < 32+2; i++ {
		newNodeID := minDistance(distance, cluster, 32+2)
		cluster[newNodeID] = true

		// Update dist[v] only if is not in cluster, there is an edge from
		// newNodeID to v, and total weight of path from src to  v through newNodeID is
		// smaller than current value of dist[v]
		for v := 0; v < 32+2; v++ {

			if !cluster[v] && Game.AdjacencyMatrix[z][newNodeID][v] != -1 && distance[newNodeID] != int(^uint(0)>>1) && distance[newNodeID]+Game.AdjacencyMatrix[z][newNodeID][v] < distance[v] {
				distance[v] = distance[newNodeID] + Game.AdjacencyMatrix[z][newNodeID][v]
				Parentarray[z][v] = newNodeID
			}
		}
	}
	//send distance of[n+1]
	//find parent of node[i]
	//and return that information to
	var botNextmove dtypes.Position
	minimumDistance := distance[33]
	markPath(z, 33)
	for i := 0; i < 32+2; i++ {
		if Parentarray[z][i] == 32 && Path[z][i] == true {
			botNextmove = Allnodes[z][i].Location
		}
	}
	currentPosition := Allnodes[z][32].Location
	updatedPosition := nextposition(currentPosition, botNextmove, step)

	//printPath(z,33)

	//fmt.Println("distance :: ",minimumDistance)
	log.Println("Completed dijkstra for channel", z)
	channel <- channels.Data{updatedPosition, minimumDistance}
	log.Println("Completed dijkstra for channel and put to channel", z, minimumDistance)
}
