package coords

import (
	"github.com/IITH-SBJoshi/concurrency-3/src/dijkstra"
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
)

// var platform [8]dtypes.Rect
// var ladder [6]dtypes.Rect
// var gem [4]dtypes.Gem

// Freepositions represents the free positions on which
// gems can land
var Freepositions [10]dtypes.Freepos

// Platform stores the diagonally opposite coordinates of
// a platform
var Platform [8]dtypes.Rect

// Ladder stores the diagonally opposite coordinates of
// a Ladder
var Ladder [6]dtypes.Rect

// Gems stores the details of a gem
var Gems [5]dtypes.Gem

// Random pos stores the teleporting random positions
var Randompos [10]dtypes.Position

// iniializePlatforms initializes positios of all platforms,
// ladders and gems
func initializePlatforms() {
	Platform[0] = dtypes.Rect{XHi:0, YHi:45, XLo:840, YLo:75}
	Platform[1] = dtypes.Rect{XHi:800, YHi:125, XLo:1200, YLo:155}
	Platform[2] = dtypes.Rect{XHi:0, YHi:200, XLo:400, YLo:230}
	Platform[3] = dtypes.Rect{XHi:600, YHi:210, XLo:1200, YLo:240}
	Platform[4] = dtypes.Rect{XHi:0, YHi:325, XLo:800, YLo:355}
	Platform[5] = dtypes.Rect{XHi:170, YHi:450, XLo:530, YLo:480}
	Platform[6] = dtypes.Rect{XHi:800, YHi:440, XLo:1200, YLo:470}
	Platform[7] = dtypes.Rect{XHi:0, YHi:550, XLo:1200, YLo:50}

	Ladder[0] = dtypes.Rect{XHi:170, YHi:450, XLo:200, YLo:550}
	Ladder[1] = dtypes.Rect{XHi:1170, YHi:440, XLo:1200, YLo:550}
	Ladder[2] = dtypes.Rect{XHi:800, YHi:210, XLo:830, YLo:440}
	Ladder[3] = dtypes.Rect{XHi:370, YHi:45, XLo:400, YLo:200}
	Ladder[4] = dtypes.Rect{XHi:100, YHi:200, XLo:130, YLo:325}
	Ladder[5] = dtypes.Rect{XHi:400, YHi:325, XLo:430, YLo:450}

	Gems[0] = dtypes.Gem{
		Gemtype: '-',
		Value:   100,
		Pos:     dtypes.Rect{XHi:200, YHi:285, XLo:230, YLo:325},
	}
	Gems[1] = dtypes.Gem{
		Gemtype: '/',
		Value:   2,
		Pos:     dtypes.Rect{XHi:520, YHi:285, XLo:550, YLo:325},
	}
	Gems[2] = dtypes.Gem{
		Gemtype: '+',
		Value:   100,
		Pos:     dtypes.Rect{XHi:450, YHi:285, XLo:480, YLo:325},
	}
	Gems[3] = dtypes.Gem{
		Gemtype: '*',
		Value:   2,
		Pos:     dtypes.Rect{XHi:600, YHi:400, XLo:630, YLo:440},
	}

	Freepositions[0] = dtypes.Freepos{Available: true, Pos: dtypes.Rect{XHi:900, YHi:80, XLo:915, YLo:100}}
	Freepositions[1] = dtypes.Freepos{Available: true, Pos: dtypes.Rect{XHi:700, YHi:510, XLo:715, YLo:530}}
	Freepositions[2] = dtypes.Freepos{Available: true, Pos: dtypes.Rect{XHi:280, YHi:285, XLo:295, YLo:305}}
	Freepositions[3] = dtypes.Freepos{Available: true, Pos: dtypes.Rect{XHi:900, YHi:400, XLo:915, YLo:420}}
	Freepositions[4] = dtypes.Freepos{Available: true, Pos: dtypes.Rect{XHi:100, YHi:5, XLo:115, YLo:25}}

	Randompos[0] = dtypes.Position{X:715, Y:25}
	Randompos[1] = dtypes.Position{X:215, Y:180}
	Randompos[2] = dtypes.Position{X:1015, Y:190}
	Randompos[3] = dtypes.Position{X:565, Y:305}  // 550, 285
	Randompos[4] = dtypes.Position{X:25, Y:530}   // 10 , 510
	Randompos[5] = dtypes.Position{X:1115, Y:530} // 1100, 510
	Randompos[6] = dtypes.Position{X:315, Y:430}  // 300, 410
	Randompos[7] = dtypes.Position{X:1015, Y:415} // 1000, 400
	Randompos[8] = dtypes.Position{X:715, Y:190}  // 700, 170
	Randompos[9] = dtypes.Position{X:25, Y:305}   // 10, 285

}

// Initialize initializes the position of all static nodes for running dijkstra
func Initialize() {
	initializePlatforms()
	for z := 0; z < 6; z++ {
		// write all distances.
		dijkstra.Allnodes[z][0] = dijkstra.StaticNode{dtypes.Position{X:15, Y:25}, NodeID:1, XNodeID:15, YNodeID:25}
		dijkstra.Allnodes[z][1] = dijkstra.StaticNode{dtypes.Position{X:385, Y:25}, NodeID:1, XNodeID:385, YNodeID:25}
		dijkstra.Allnodes[z][2] = dijkstra.StaticNode{dtypes.Position{X:855, Y:25}, NodeID:1, XNodeID:855, YNodeID:25}
		dijkstra.Allnodes[z][3] = dijkstra.StaticNode{dtypes.Position{X:785, Y:105}, NodeID:2, XNodeID:785, YNodeID:105}
		dijkstra.Allnodes[z][4] = dijkstra.StaticNode{dtypes.Position{X:855, Y:105}, NodeID:2, XNodeID:855, YNodeID:105}
		dijkstra.Allnodes[z][5] = dijkstra.StaticNode{dtypes.Position{X:1185, Y:105}, NodeID:2, XNodeID:1185, YNodeID:105}
		dijkstra.Allnodes[z][6] = dijkstra.StaticNode{dtypes.Position{X:15, Y:180}, NodeID:3, XNodeID:15, YNodeID:180}
		dijkstra.Allnodes[z][7] = dijkstra.StaticNode{dtypes.Position{X:385, Y:180}, NodeID:3, XNodeID:385, YNodeID:180}
		dijkstra.Allnodes[z][8] = dijkstra.StaticNode{dtypes.Position{X:415, Y:180}, NodeID:3, XNodeID:415, YNodeID:180}
		dijkstra.Allnodes[z][9] = dijkstra.StaticNode{dtypes.Position{X:585, Y:190}, NodeID:4, XNodeID:585, YNodeID:190}
		dijkstra.Allnodes[z][10] = dijkstra.StaticNode{dtypes.Position{X:785, Y:190}, NodeID:4, XNodeID:785, YNodeID:190}
		dijkstra.Allnodes[z][11] = dijkstra.StaticNode{dtypes.Position{X:815, Y:190}, NodeID:4, XNodeID:815, YNodeID:190}
		dijkstra.Allnodes[z][12] = dijkstra.StaticNode{dtypes.Position{X:1185, Y:190}, NodeID:4, XNodeID:1185, YNodeID:190}
		dijkstra.Allnodes[z][13] = dijkstra.StaticNode{dtypes.Position{X:15, Y:305}, NodeID:5, XNodeID:15, YNodeID:305}
		dijkstra.Allnodes[z][14] = dijkstra.StaticNode{dtypes.Position{X:415, Y:305}, NodeID:5, XNodeID:415, YNodeID:305}
		dijkstra.Allnodes[z][15] = dijkstra.StaticNode{dtypes.Position{X:585, Y:305}, NodeID:5, XNodeID:585, YNodeID:305}
		dijkstra.Allnodes[z][16] = dijkstra.StaticNode{dtypes.Position{X:815, Y:305}, NodeID:5, XNodeID:815, YNodeID:305}
		dijkstra.Allnodes[z][17] = dijkstra.StaticNode{dtypes.Position{X:155, Y:430}, NodeID:6, XNodeID:155, YNodeID:430}
		dijkstra.Allnodes[z][18] = dijkstra.StaticNode{dtypes.Position{X:185, Y:430}, NodeID:6, XNodeID:185, YNodeID:430}
		dijkstra.Allnodes[z][19] = dijkstra.StaticNode{dtypes.Position{X:415, Y:430}, NodeID:6, XNodeID:415, YNodeID:430}
		dijkstra.Allnodes[z][20] = dijkstra.StaticNode{dtypes.Position{X:545, Y:430}, NodeID:6, XNodeID:545, YNodeID:430}
		dijkstra.Allnodes[z][21] = dijkstra.StaticNode{dtypes.Position{X:785, Y:420}, NodeID:7, XNodeID:785, YNodeID:440}
		dijkstra.Allnodes[z][22] = dijkstra.StaticNode{dtypes.Position{X:815, Y:420}, NodeID:7, XNodeID:815, YNodeID:440}
		dijkstra.Allnodes[z][23] = dijkstra.StaticNode{dtypes.Position{X:1185, Y:420}, NodeID:7, XNodeID:1185, YNodeID:440}
		dijkstra.Allnodes[z][24] = dijkstra.StaticNode{dtypes.Position{X:15, Y:530}, NodeID:8, XNodeID:15, YNodeID:530}
		dijkstra.Allnodes[z][25] = dijkstra.StaticNode{dtypes.Position{X:155, Y:530}, NodeID:8, XNodeID:155, YNodeID:530}
		dijkstra.Allnodes[z][26] = dijkstra.StaticNode{dtypes.Position{X:185, Y:530}, NodeID:8, XNodeID:185, YNodeID:530}
		dijkstra.Allnodes[z][27] = dijkstra.StaticNode{dtypes.Position{X:545, Y:530}, NodeID:8, XNodeID:545, YNodeID:530}
		dijkstra.Allnodes[z][28] = dijkstra.StaticNode{dtypes.Position{X:785, Y:530}, NodeID:8, XNodeID:785, YNodeID:530}
		dijkstra.Allnodes[z][29] = dijkstra.StaticNode{dtypes.Position{X:1185, Y:530}, NodeID:8, XNodeID:1185, YNodeID:530}
		dijkstra.Allnodes[z][30] = dijkstra.StaticNode{dtypes.Position{X:115, Y:180}, NodeID:3, XNodeID:115, YNodeID:180}
		dijkstra.Allnodes[z][31] = dijkstra.StaticNode{dtypes.Position{X:115, Y:305}, NodeID:32, XNodeID:115, YNodeID:305}

		//dijkstra.Allnodes[z][]={{,},,,}
		// if there is no edge between node i and j then value corresponding them will be -1
		for i := 0; i < 34; i++ {
			for j := 0; j < 34; j++ {
				dijkstra.Game.AdjacencyMatrix[z][i][j] = -1
			}
		}

		// initialize global adjacency matrix.
		dijkstra.Game.AdjacencyMatrix[z][0][1] = 370
		dijkstra.Game.AdjacencyMatrix[z][1][0] = 370
		dijkstra.Game.AdjacencyMatrix[z][1][2] = 470
		dijkstra.Game.AdjacencyMatrix[z][2][1] = 470
		dijkstra.Game.AdjacencyMatrix[z][1][7] = 155
		dijkstra.Game.AdjacencyMatrix[z][7][1] = 155
		dijkstra.Game.AdjacencyMatrix[z][2][4] = 80
		dijkstra.Game.AdjacencyMatrix[z][3][4] = 70
		dijkstra.Game.AdjacencyMatrix[z][4][3] = 70
		dijkstra.Game.AdjacencyMatrix[z][4][5] = 330
		dijkstra.Game.AdjacencyMatrix[z][5][4] = 330
		dijkstra.Game.AdjacencyMatrix[z][6][30] = 100
		dijkstra.Game.AdjacencyMatrix[z][30][6] = 100
		dijkstra.Game.AdjacencyMatrix[z][7][30] = 270
		dijkstra.Game.AdjacencyMatrix[z][30][7] = 270
		dijkstra.Game.AdjacencyMatrix[z][7][8] = 30
		dijkstra.Game.AdjacencyMatrix[z][8][7] = 30
		dijkstra.Game.AdjacencyMatrix[z][3][10] = 85
		dijkstra.Game.AdjacencyMatrix[z][30][31] = 125
		dijkstra.Game.AdjacencyMatrix[z][31][30] = 125
		dijkstra.Game.AdjacencyMatrix[z][13][31] = 100
		dijkstra.Game.AdjacencyMatrix[z][31][13] = 100
		dijkstra.Game.AdjacencyMatrix[z][31][14] = 300
		dijkstra.Game.AdjacencyMatrix[z][14][31] = 300
		dijkstra.Game.AdjacencyMatrix[z][14][15] = 170
		dijkstra.Game.AdjacencyMatrix[z][15][14] = 170
		dijkstra.Game.AdjacencyMatrix[z][15][16] = 230
		dijkstra.Game.AdjacencyMatrix[z][17][15] = 230
		dijkstra.Game.AdjacencyMatrix[z][9][10] = 200
		dijkstra.Game.AdjacencyMatrix[z][10][9] = 200
		dijkstra.Game.AdjacencyMatrix[z][10][11] = 30
		dijkstra.Game.AdjacencyMatrix[z][11][10] = 30
		dijkstra.Game.AdjacencyMatrix[z][11][12] = 370
		dijkstra.Game.AdjacencyMatrix[z][12][11] = 370
		dijkstra.Game.AdjacencyMatrix[z][9][15] = 115
		dijkstra.Game.AdjacencyMatrix[z][8][14] = 125
		dijkstra.Game.AdjacencyMatrix[z][11][16] = 115
		dijkstra.Game.AdjacencyMatrix[z][16][11] = 115
		dijkstra.Game.AdjacencyMatrix[z][16][22] = 115
		dijkstra.Game.AdjacencyMatrix[z][22][16] = 115
		dijkstra.Game.AdjacencyMatrix[z][14][19] = 125
		dijkstra.Game.AdjacencyMatrix[z][19][14] = 125
		dijkstra.Game.AdjacencyMatrix[z][17][25] = 100
		dijkstra.Game.AdjacencyMatrix[z][18][26] = 100
		dijkstra.Game.AdjacencyMatrix[z][26][18] = 100
		dijkstra.Game.AdjacencyMatrix[z][20][27] = 100
		dijkstra.Game.AdjacencyMatrix[z][17][18] = 30
		dijkstra.Game.AdjacencyMatrix[z][18][17] = 30
		dijkstra.Game.AdjacencyMatrix[z][18][19] = 230
		dijkstra.Game.AdjacencyMatrix[z][19][18] = 230
		dijkstra.Game.AdjacencyMatrix[z][19][20] = 130
		dijkstra.Game.AdjacencyMatrix[z][20][19] = 130
		dijkstra.Game.AdjacencyMatrix[z][21][22] = 30
		dijkstra.Game.AdjacencyMatrix[z][22][21] = 30
		dijkstra.Game.AdjacencyMatrix[z][22][23] = 370
		dijkstra.Game.AdjacencyMatrix[z][23][22] = 370
		dijkstra.Game.AdjacencyMatrix[z][21][28] = 110
		dijkstra.Game.AdjacencyMatrix[z][23][29] = 110
		dijkstra.Game.AdjacencyMatrix[z][29][23] = 110
		dijkstra.Game.AdjacencyMatrix[z][24][25] = 100
		dijkstra.Game.AdjacencyMatrix[z][25][24] = 100
		dijkstra.Game.AdjacencyMatrix[z][25][26] = 30
		dijkstra.Game.AdjacencyMatrix[z][26][25] = 30
		dijkstra.Game.AdjacencyMatrix[z][26][27] = 360
		dijkstra.Game.AdjacencyMatrix[z][27][26] = 360
		dijkstra.Game.AdjacencyMatrix[z][27][28] = 240
		dijkstra.Game.AdjacencyMatrix[z][28][27] = 240
		dijkstra.Game.AdjacencyMatrix[z][28][29] = 400
		dijkstra.Game.AdjacencyMatrix[z][29][28] = 400
		//initialization of parent
		for i := 0; i < 32+2; i++ {
			dijkstra.Parentarray[z][i] = -1
		}
	}
}
