package play_node_runner

import (
	"log"
	"sync"
	"time"

	"github.com/IITH-SBJoshi/concurrency-3/src/channels"
	"github.com/IITH-SBJoshi/concurrency-3/src/client"
	"github.com/IITH-SBJoshi/concurrency-3/src/coords"
	"github.com/IITH-SBJoshi/concurrency-3/src/dijkstra"
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
	handler "github.com/IITH-SBJoshi/concurrency-3/src/handlers"
	"github.com/IITH-SBJoshi/concurrency-3/src/health"
	"github.com/IITH-SBJoshi/concurrency-3/src/locs"
	"github.com/IITH-SBJoshi/concurrency-3/src/platform"

	"github.com/gorilla/websocket"
)

var lock sync.Mutex

// regularUpdater regularly asks for update to clients
func regularUpdater(conn *websocket.Conn, requestChannelClient, receiveChannelClient chan dtypes.Event, id int) {

	ticker := time.NewTicker(30 * time.Millisecond)
	go func() {
		for range ticker.C {
			var event dtypes.Event
			event = dtypes.Event{
				EventType: "SendUpdate",
			}
			lock.Lock()
			err := conn.WriteJSON(event)
			lock.Unlock()
			if err != nil {
				log.Println(id, "Error writing json in sendupdate", err)
			}
			log.Println(id, "written sendupdate json", event.GetStr())
			// }

			lock.Lock()
			err = conn.ReadJSON(&event)
			lock.Unlock()
			log.Println(id, " read json from regular updater", event.GetStr())
			if err != nil {
				log.Println(id, " Error reading jsonn message.", err)
			}
			log.Println(id, " making request to request channel client")

			requestChannelClient <- event
			log.Println(id, " event written to requestChannelClient", event.GetStr())
		}
	}()
}

// sendResponse sends respone to a client
func sendResponse(receiveChannelClient chan dtypes.Event, conn *websocket.Conn, id int) {
	responseMsg := <-receiveChannelClient
	lock.Lock()
	err := conn.WriteJSON(responseMsg)
	lock.Unlock()
	if err != nil {
		log.Println(id, " Error writing json.", err)
	}
	log.Println(id, " written json from send response")
}

// PlayNodeRunner is the event loop of NodeRunner
func PlayNodeRunner(requestChannelServer, firstRespondChannelServer, secondRespondChannelServer chan dtypes.Event, gameWinChannel chan int, firstClient, secondClient *client.Client) {
	coords.Initialize()
	platform.Initialize()
	locs.InitializeLocations()
	channels.ChannelInitialization()
	handler.SetGameWinChannel(gameWinChannel)
	health.SetHealth(1000)
	health.SetDecayParams(10, 500)

	go health.DecayPlayer1()
	go health.DecayPlayer2()
	go regularUpdater(firstClient.GetWSocket(), firstClient.GetRequestChannel(), firstClient.GetReceiveChannel(), 0)
	go regularUpdater(secondClient.GetWSocket(), secondClient.GetRequestChannel(), secondClient.GetReceiveChannel(), 1)
	// go serverComputations(firstClient.GetRequestChannel(), secondClient.GetRequestChannel(), firstRespondChannelServer, secondRespondChannelServer, requestChannelServer)
	go serverComputations(firstClient.GetRequestChannel(), secondClient.GetRequestChannel(), firstClient.GetWSocket(), secondClient.GetWSocket(), requestChannelServer)
	go serverReceiveComputations(firstClient.GetRequestChannel(), secondClient.GetRequestChannel(), firstRespondChannelServer, secondRespondChannelServer, requestChannelServer)
}

// readConnections reads the websocket connections
func readConnections(conn *websocket.Conn, requestChannel chan dtypes.Event, id int) {
	for {
		event := dtypes.Event{}
		lock.Lock()
		err := conn.ReadJSON(&event)
		lock.Unlock()
		log.Println(id, " read json from readConnections", event.GetStr())
		if err != nil {
			log.Println(id, " Error reading json.", err)
		}
		requestChannel <- event
	}
}

// serverReceiveComputations sends comptations randomly from request channels
// of clients to request channel of server. the select block selects randomly
// from the two player channels thus ensuring fairness among the players
func serverReceiveComputations(firstClientRequestChannel, secondClientRequestChannel, firstRespondChannelServer, secondRespondChannelServer, requestChannelServer chan dtypes.Event) {
	log.Println("started running servercomputations")
	for {
		select {
		case latestState := <-firstClientRequestChannel:
			log.Println("first client request chanel passed the info to server")
			requestChannelServer <- latestState

		case latestState := <-secondClientRequestChannel:
			requestChannelServer <- latestState
			log.Println("second client request channel passed the info to server")

		}
	}
}

// serverComputations handles computations from a server
func serverComputations(firstClientRequestChannel, secondClientRequestChannel chan dtypes.Event, firstConn, secondConn *websocket.Conn, requestChannelServer chan dtypes.Event) {

	for {
		log.Println("Server received a event msg to compute at Server computations")
		playerPositions := <-requestChannelServer
		playerPositions = locs.GetCurrentLocations(playerPositions)
		log.Println("Got current player positions:", playerPositions.GetStr())
		updatedPlayerPositions := handler.Handle(playerPositions)
		log.Println("updated positions of players ", updatedPlayerPositions.GetStr())
		// updatedPlayerPositions := playerPositions
		updatedBotPositions := dijkstra.UpdateBots(updatedPlayerPositions)
		locs.SetCurrentLocations(updatedBotPositions)
		log.Println("upadted positions of bots", updatedBotPositions.GetStr())
		// updatedBotPositions := updatedPlayerPositions
		lock.Lock()
		err := firstConn.WriteJSON(updatedBotPositions)
		lock.Unlock()
		log.Println(" written update json")
		if err != nil {
			log.Println(" Error writing json.", err)
		}
		log.Println("Updated bot positon sent to first client socket")
		// secondRespondChannelServer <- updatedBotPositions
		lock.Lock()
		err = secondConn.WriteJSON(updatedBotPositions)
		lock.Unlock()
		log.Println(" written update json")
		if err != nil {
			log.Println(" Error writing json.", err)
		}
		log.Println("Updated bot positions sent to second client socket")
		// log.Println("updated bot positions sent to secondrespond channelserver")
		log.Println("Updated bot positons send to client objects")
	}
}
