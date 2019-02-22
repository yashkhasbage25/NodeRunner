package dijkstra

import (

	"dtypes"
	"fmt"
)

// staticNode number of static nodes are 32 depended on construction of graph ..here manually measured
type staticNode struct {
	location dtypes.Position
	nodeID int 
	xnodeID int
	ynodeID int
	//type of node ..ladder or ..end of path or ..
}

type matrix struct { // 32 is number of static node and we are adding 2 dynamic node.
	adjacencyMatrix [6][32+2][32+2]int
	size int; // 34
	
}

// global game matrix 
var game matrix
var parentarray [6][32+2] int
var allNodes [6][32+2] staticNode
func  initialize() {
	for z:=0;z<6;z++{
		// write all distances.
		allNodes[z][0]=staticNode{dtypes.Position{15,25},1,15,25}
		allNodes[z][1]=staticNode{dtypes.Position{385,25},1,385,25}
		allNodes[z][2]=staticNode{dtypes.Position{855,25},1,855,25}
		allNodes[z][3]=staticNode{dtypes.Position{785,105},2,785,105}
		allNodes[z][4]=staticNode{dtypes.Position{855,105},2,855,105}
		allNodes[z][5]=staticNode{dtypes.Position{1185,105},2,1185,105}
		allNodes[z][6]=staticNode{dtypes.Position{15,180},3,15,180}
		allNodes[z][7]=staticNode{dtypes.Position{385,180},3,385,180}
		allNodes[z][8]=staticNode{dtypes.Position{415,180},3,415,180}
		allNodes[z][9]=staticNode{dtypes.Position{585,190},4,585,190}
		allNodes[z][10]=staticNode{dtypes.Position{785,190},4,785,190}
		allNodes[z][11]=staticNode{dtypes.Position{815,190},4,815,190}
		allNodes[z][12]=staticNode{dtypes.Position{1185,190},4,1185,190}
		allNodes[z][13]=staticNode{dtypes.Position{15,305},5,15,305}
		allNodes[z][14]=staticNode{dtypes.Position{415,305},5,415,305}
		allNodes[z][15]=staticNode{dtypes.Position{585,305},5,585,305}
		allNodes[z][16]=staticNode{dtypes.Position{815,305},5,815,305}
		allNodes[z][17]=staticNode{dtypes.Position{155,430},6,155,430}
		allNodes[z][18]=staticNode{dtypes.Position{185,430},6,185,430}
		allNodes[z][19]=staticNode{dtypes.Position{415,430},6,415,430}
		allNodes[z][20]=staticNode{dtypes.Position{545,430},6,545,430}
		allNodes[z][21]=staticNode{dtypes.Position{785,420},7,785,440}
		allNodes[z][22]=staticNode{dtypes.Position{815,420},7,815,440}
		allNodes[z][23]=staticNode{dtypes.Position{1185,420},7,1185,440}
		allNodes[z][24]=staticNode{dtypes.Position{15,530},8,15,530}
		allNodes[z][25]=staticNode{dtypes.Position{155,530},8,155,530}
		allNodes[z][26]=staticNode{dtypes.Position{185,530},8,185,530}
		allNodes[z][27]=staticNode{dtypes.Position{545,530},8,545,530}
		allNodes[z][28]=staticNode{dtypes.Position{785,530},8,785,530}
		allNodes[z][29]=staticNode{dtypes.Position{1185,530},8,1185,530}
		allNodes[z][30]=staticNode{dtypes.Position{115,180},3,115,180}
		allNodes[z][31]=staticNode{dtypes.Position{115,305},32,115,305}


		//allNodes[z][]={{,},,,}

		// initialize global adjacency matrix.
		 game.adjacencyMatrix[z][0][1]=370
		 game.adjacencyMatrix[z][1][0]=370
		 game.adjacencyMatrix[z][1][2]=470
		 game.adjacencyMatrix[z][2][1]=470
		 game.adjacencyMatrix[z][1][7]=155
		 game.adjacencyMatrix[z][7][1]=155
		 game.adjacencyMatrix[z][2][4]=80
		 game.adjacencyMatrix[z][3][4]=70
		 game.adjacencyMatrix[z][4][3]=70
		 game.adjacencyMatrix[z][4][5]=330
		 game.adjacencyMatrix[z][5][4]=330
		 game.adjacencyMatrix[z][6][30]=100
		 game.adjacencyMatrix[z][30][6]=100
		 game.adjacencyMatrix[z][7][30]=270
		 game.adjacencyMatrix[z][30][7]=270
		 game.adjacencyMatrix[z][7][8]=30
		 game.adjacencyMatrix[z][8][7]=30
		 game.adjacencyMatrix[z][3][10]=85
		 game.adjacencyMatrix[z][30][31]=125
		 game.adjacencyMatrix[z][31][30]=125
		 game.adjacencyMatrix[z][13][31]=100
		 game.adjacencyMatrix[z][31][13]=100
		 game.adjacencyMatrix[z][31][14]=300
		 game.adjacencyMatrix[z][14][31]=300
		 game.adjacencyMatrix[z][14][15]=170
		 game.adjacencyMatrix[z][15][14]=170
		 game.adjacencyMatrix[z][15][16]=230
		 game.adjacencyMatrix[z][17][15]=230
		 game.adjacencyMatrix[z][9][10]=200
		 game.adjacencyMatrix[z][10][9]=200
		 game.adjacencyMatrix[z][10][11]=30
		 game.adjacencyMatrix[z][11][10]=30
		 game.adjacencyMatrix[z][11][12]=370
		 game.adjacencyMatrix[z][12][11]=370
		 game.adjacencyMatrix[z][9][15]=115
		 game.adjacencyMatrix[z][8][14]=125
		 game.adjacencyMatrix[z][11][16]=115
		 game.adjacencyMatrix[z][16][11]=115
		 game.adjacencyMatrix[z][16][22]=115
		 game.adjacencyMatrix[z][22][16]=115
		 game.adjacencyMatrix[z][14][19]=125
		 game.adjacencyMatrix[z][19][14]=125
		 game.adjacencyMatrix[z][17][25]=100
		 game.adjacencyMatrix[z][18][26]=100
		 game.adjacencyMatrix[z][26][18]=100
		 game.adjacencyMatrix[z][20][27]=100
		 game.adjacencyMatrix[z][17][18]=30
		 game.adjacencyMatrix[z][18][17]=30
		 game.adjacencyMatrix[z][18][19]=230
		 game.adjacencyMatrix[z][19][18]=230
		 game.adjacencyMatrix[z][19][20]=130
		 game.adjacencyMatrix[z][20][19]=130
		 game.adjacencyMatrix[z][21][22]=30
		 game.adjacencyMatrix[z][22][21]=30
		 game.adjacencyMatrix[z][22][23]=370
		 game.adjacencyMatrix[z][23][22]=370
		 game.adjacencyMatrix[z][21][28]=110
		 game.adjacencyMatrix[z][23][29]=110
		 game.adjacencyMatrix[z][29][23]=110
		 game.adjacencyMatrix[z][24][25]=100
		 game.adjacencyMatrix[z][25][24]=100
		 game.adjacencyMatrix[z][25][26]=30
		 game.adjacencyMatrix[z][26][25]=30
		 game.adjacencyMatrix[z][26][27]=360
		 game.adjacencyMatrix[z][27][26]=360
		 game.adjacencyMatrix[z][27][28]=240
		 game.adjacencyMatrix[z][28][27]=240
		 game.adjacencyMatrix[z][28][29]=400
		 game.adjacencyMatrix[z][29][28]=400
		 	//initialization of parent
		for i:=0;i<32+2;i++{
			parentarray[z][i]=-1
		} 
		
	}

}
func Abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}
func setter(i int ,n int, z int,flag bool){
	if flag==true{
		game.adjacencyMatrix[z][n][i]=Abs(allNodes[z][i].location.Y-allNodes[z][n].location.Y)
		game.adjacencyMatrix[z][i][n]=Abs(allNodes[z][i].location.Y-allNodes[z][n].location.Y)
	}	else  {
		game.adjacencyMatrix[z][n][i]=Abs(allNodes[z][i].location.X-allNodes[z][n].location.X)
		game.adjacencyMatrix[z][i][n]=Abs(allNodes[z][i].location.X-allNodes[z][n].location.X)
	}
	
}

func addDynamicnode(bot dtypes.Position,player dtypes.Position,z int) {
	allNodes[z][32]= staticNode{dtypes.Position{bot.X,bot.Y},9,bot.X,bot.Y} //added bot at position equal to n.
	allNodes[z][32+1]=staticNode{dtypes.Position{player.X,player.Y},10 ,player.X, player.Y}//added player
	for i:=0; i<32 ; i++ { 
		if allNodes[z][i].xnodeID==allNodes[z][32].xnodeID {
			setter(i,32,z,true)
		}	else if allNodes[z][i].ynodeID==allNodes[z][32].ynodeID {
			setter(i,32,z,false)
		}
	}
	for i:=0; i<32+1; i++ {
		if allNodes[z][i].xnodeID==allNodes[z][32+1].xnodeID {
			setter(i,32+1,z,true)
		}	else if allNodes[z][i].ynodeID==allNodes[z][32+1].ynodeID {
			setter(i,32+1,z,false)
		}
	}
	
}
func minDistance(distance []int, cluster []bool, size int) int {
	
	 var min_index int;
	 min := int(^uint(0)>> 1)  
     for v:=0; v<size;v++{  // we have to impliment this parallaly
     	if cluster[v]==false&&distance[v]<=min {
     		min= distance[v]
     		min_index=v
     	}
     } 
     return min_index
}
func printPath(distance []int,z int,node int)  {

	if(parentarray[z][node]!=-1) {
		if parentarray[z][node]!=32{
			printPath(distance,z,parentarray[z][node])
		}
		fmt.Println(parentarray[z][node])
	}
}
func RunDijkstra(bot dtypes.Position,player dtypes.Position, z int)(dtypes.Position,int) {

	initialize()
	addDynamicnode(bot,player,z)
	//we have source as node with nodeid
	//var distance [32+2] int
	distance := make([]int, 34)
	cluster  := make([]bool, 34) 
	//var cluster [32+2] bool // cluster[i] will be true if node i is included in shortest 
                          // path tree or shortest distance from src to i is finalize

	for i :=0;i<32+2;i++{
        distance[i] =int(^uint(0)>> 1)// it is int_max in golang
        cluster[i]=false
    }
    distance[32]=0  // distance of src from src is 0

    for i:=0;i<32+2;i++{
    	newnodeID:=minDistance(distance, cluster,32+2)
    	cluster[newnodeID]=true
    	
   
         // Update dist[v] only if is not in cluster, there is an edge from  
         // newnodeID to v, and total weight of path from src to  v through newnodeID is  
         // smaller than current value of dist[v] 
         for v:=0; v < 32+2; v++ {

         	if  !cluster[v] && game.adjacencyMatrix[z][newnodeID][v]!=0 && distance[newnodeID] != int(^uint(0)>> 1)&& distance[newnodeID]+game.adjacencyMatrix[z][newnodeID][v]  < distance[v] {
         			distance[v] = distance[newnodeID] + game.adjacencyMatrix[z][newnodeID][v]
         			parentarray[z][v]=newnodeID
         		}    
         } 
    }
    //send distance of[n+1]
    //find parent of node[i]
    //and return that information to 
    var botNextmove dtypes.Position
    minimumDistance:=distance[33]
    for i:=0;i<32+2;i++{
    	if parentarray[z][i]==32{
    		botNextmove=allNodes[z][i].location
    	}
    }
    printPath(distance,z,33)

    fmt.Println("distance :: ",minimumDistance)
    return botNextmove,minimumDistance
}