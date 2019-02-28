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
	Platform[0] = dtypes.Rect{0, 45, 840, 75}
	Platform[1] = dtypes.Rect{800, 125, 1200, 155}
	Platform[2] = dtypes.Rect{0, 200, 400, 230}
	Platform[3] = dtypes.Rect{600, 210, 1200, 240}
	Platform[4] = dtypes.Rect{0, 325, 800, 355}
	Platform[5] = dtypes.Rect{170, 450, 530, 480}
	Platform[6] = dtypes.Rect{800, 440, 1200, 470}
	Platform[7] = dtypes.Rect{0, 550, 1200, 50}

	Ladder[0] = dtypes.Rect{170, 450, 200, 550}
	Ladder[1] = dtypes.Rect{1170, 440, 1200, 550}
	Ladder[2] = dtypes.Rect{800, 210, 830, 440}
	Ladder[3] = dtypes.Rect{370, 45, 400, 200}
	Ladder[4] = dtypes.Rect{100, 200, 130, 325}
	Ladder[5] = dtypes.Rect{400, 325, 430, 450}

	Gems[0] = dtypes.Gem{
		Gemtype: '-',
		Value:   100,
		Pos:     dtypes.Rect{200, 285, 230, 325},
	}
	Gems[1] = dtypes.Gem{
		Gemtype: '/',
		Value:   2,
		Pos:     dtypes.Rect{520, 285, 550, 325},
	}
	Gems[2] = dtypes.Gem{
		Gemtype: '+',
		Value:   100,
		Pos:     dtypes.Rect{450, 285, 480, 325},
	}
	Gems[3] = dtypes.Gem{
		Gemtype: '*',
		Value:   2,
		Pos:     dtypes.Rect{600, 510, 630, 550},
	}

	Freepositions[0] = dtypes.Freepos{Available: true, Pos: dtypes.Rect{900, 80, 930, 120}}
	Freepositions[1] = dtypes.Freepos{Available: true, Pos: dtypes.Rect{700, 510, 730, 550}}
	Freepositions[2] = dtypes.Freepos{Available: true, Pos: dtypes.Rect{280, 285, 310, 325}}
	Freepositions[3] = dtypes.Freepos{Available: true, Pos: dtypes.Rect{900, 400, 930, 440}}

	Randompos[0] = dtypes.Position{715, 25}
	Randompos[1] = dtypes.Position{215, 180}
	Randompos[2] = dtypes.Position{1015, 190}
	Randompos[3] = dtypes.Position{565, 305}  // 550, 285
	Randompos[4] = dtypes.Position{25, 530}   // 10 , 510
	Randompos[5] = dtypes.Position{1115, 530} // 1100, 510
	Randompos[6] = dtypes.Position{315, 430}  // 300, 410
	Randompos[7] = dtypes.Position{1015, 420} // 1000, 400
	Randompos[8] = dtypes.Position{715, 190}  // 700, 170
	Randompos[9] = dtypes.Position{25, 305}   // 10, 285

}

// Initialize initializes the position of all static nodes for running dijkstra
func Initialize() {
	initializePlatforms()
	for z := 0; z < 6; z++ {
		// write all distances.
		dijkstra.Allnodes[z][0] = dijkstra.StaticNode{dtypes.Position{15, 25}, 1, 15, 25}
		dijkstra.Allnodes[z][1] = dijkstra.StaticNode{dtypes.Position{385, 25}, 1, 385, 25}
		dijkstra.Allnodes[z][2] = dijkstra.StaticNode{dtypes.Position{855, 25}, 1, 855, 25}
		dijkstra.Allnodes[z][3] = dijkstra.StaticNode{dtypes.Position{785, 105}, 2, 785, 105}
		dijkstra.Allnodes[z][4] = dijkstra.StaticNode{dtypes.Position{855, 105}, 2, 855, 105}
		dijkstra.Allnodes[z][5] = dijkstra.StaticNode{dtypes.Position{1185, 105}, 2, 1185, 105}
		dijkstra.Allnodes[z][6] = dijkstra.StaticNode{dtypes.Position{15, 180}, 3, 15, 180}
		dijkstra.Allnodes[z][7] = dijkstra.StaticNode{dtypes.Position{385, 180}, 3, 385, 180}
		dijkstra.Allnodes[z][8] = dijkstra.StaticNode{dtypes.Position{415, 180}, 3, 415, 180}
		dijkstra.Allnodes[z][9] = dijkstra.StaticNode{dtypes.Position{585, 190}, 4, 585, 190}
		dijkstra.Allnodes[z][10] = dijkstra.StaticNode{dtypes.Position{785, 190}, 4, 785, 190}
		dijkstra.Allnodes[z][11] = dijkstra.StaticNode{dtypes.Position{815, 190}, 4, 815, 190}
		dijkstra.Allnodes[z][12] = dijkstra.StaticNode{dtypes.Position{1185, 190}, 4, 1185, 190}
		dijkstra.Allnodes[z][13] = dijkstra.StaticNode{dtypes.Position{15, 305}, 5, 15, 305}
		dijkstra.Allnodes[z][14] = dijkstra.StaticNode{dtypes.Position{415, 305}, 5, 415, 305}
		dijkstra.Allnodes[z][15] = dijkstra.StaticNode{dtypes.Position{585, 305}, 5, 585, 305}
		dijkstra.Allnodes[z][16] = dijkstra.StaticNode{dtypes.Position{815, 305}, 5, 815, 305}
		dijkstra.Allnodes[z][17] = dijkstra.StaticNode{dtypes.Position{155, 430}, 6, 155, 430}
		dijkstra.Allnodes[z][18] = dijkstra.StaticNode{dtypes.Position{185, 430}, 6, 185, 430}
		dijkstra.Allnodes[z][19] = dijkstra.StaticNode{dtypes.Position{415, 430}, 6, 415, 430}
		dijkstra.Allnodes[z][20] = dijkstra.StaticNode{dtypes.Position{545, 430}, 6, 545, 430}
		dijkstra.Allnodes[z][21] = dijkstra.StaticNode{dtypes.Position{785, 420}, 7, 785, 420}
		dijkstra.Allnodes[z][22] = dijkstra.StaticNode{dtypes.Position{815, 420}, 7, 815, 420}
		dijkstra.Allnodes[z][23] = dijkstra.StaticNode{dtypes.Position{1185, 420}, 7, 1185, 420}
		dijkstra.Allnodes[z][24] = dijkstra.StaticNode{dtypes.Position{15, 530}, 8, 15, 530}
		dijkstra.Allnodes[z][25] = dijkstra.StaticNode{dtypes.Position{155, 530}, 8, 155, 530}
		dijkstra.Allnodes[z][26] = dijkstra.StaticNode{dtypes.Position{185, 530}, 8, 185, 530}
		dijkstra.Allnodes[z][27] = dijkstra.StaticNode{dtypes.Position{545, 530}, 8, 545, 530}
		dijkstra.Allnodes[z][28] = dijkstra.StaticNode{dtypes.Position{785, 530}, 8, 785, 530}
		dijkstra.Allnodes[z][29] = dijkstra.StaticNode{dtypes.Position{1185, 530}, 8, 1185, 530}
		dijkstra.Allnodes[z][30] = dijkstra.StaticNode{dtypes.Position{115, 180}, 3, 115, 180}
		dijkstra.Allnodes[z][31] = dijkstra.StaticNode{dtypes.Position{115, 305}, 32, 115, 305}

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
