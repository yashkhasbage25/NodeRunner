package play_node_runner

import (
	"log"
	"time"

	"github.com/IITH-SBJoshi/concurrency-3/src/client"
	"github.com/IITH-SBJoshi/concurrency-3/src/coords"
	"github.com/IITH-SBJoshi/concurrency-3/src/dijkstra"
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
	handler "github.com/IITH-SBJoshi/concurrency-3/src/handlers"
	"github.com/IITH-SBJoshi/concurrency-3/src/health"
	"github.com/IITH-SBJoshi/concurrency-3/src/platform"

	"github.com/gorilla/websocket"
)

func regularUpdater(conn *websocket.Conn, requestChannelClient chan dtypes.Event) {
	var event dtypes.Event
	ticker := time.NewTicker(1000 * time.Millisecond)
	go func() {
		for range ticker.C {
			event = dtypes.Event{
				EventType: "SendUpdate",
			}

			err := conn.WriteJSON(event)
			log.Println("written sendupdate json")
			if err != nil {
				log.Println("Error writing json.", err)
			}
			// err = json.Unmarshal([]byte(), v)

			// msgType, msg, err := conn.ReadMessage()
			// if err != nil {
			//
			// }
			err = conn.ReadJSON(&event)
			// r := "                                                                                                                                                                          "
			// err = json.Unmarshal([]byte(r), &event)
			log.Println("read json from regular updater", event.GetStr())
			if err != nil {
				log.Println("Error reading jsonn message.", err)
			}
			requestChannelClient <- event
		}
	}()
}

func sendResponse(receiveChannelClient chan dtypes.Event, conn *websocket.Conn) {
	responseMsg := <-receiveChannelClient
	err := conn.WriteJSON(responseMsg)
	if err != nil {
		log.Println("Error writing json.", err)
	}
	log.Println("written json from send response")
}

// PlayNodeRunner is the event loop of NodeRunner
func PlayNodeRunner(requestChannelServer, firstRespondChannelServer, secondRespondChannelServer chan dtypes.Event, gameWinChannel chan int, firstClient, secondClient *client.Client) {
	coords.Initialize()
	platform.Initialize()
	handler.SetGameWinChannel(gameWinChannel)
	health.SetHealth("p1", 1000)
	health.SetHealth("p2", 1000)
	health.SetDecayParams(1, 500)

	go health.DecayPlayer1()
	go health.DecayPlayer2()
	go regularUpdater(firstClient.GetWSocket(), firstClient.GetRequestChannel())
	go regularUpdater(secondClient.GetWSocket(), secondClient.GetRequestChannel())
	go serverComputations(firstClient.GetRequestChannel(), secondClient.GetRequestChannel(), firstRespondChannelServer, secondRespondChannelServer, requestChannelServer)
	go serverReceiveComputations(firstClient.GetRequestChannel(), secondClient.GetRequestChannel(), firstRespondChannelServer, secondRespondChannelServer, requestChannelServer)
	go sendResponse(firstClient.GetReceiveChannel(), firstClient.GetWSocket())
	go sendResponse(secondClient.GetReceiveChannel(), secondClient.GetWSocket())
	go readConnections(firstClient.GetWSocket(), firstClient.GetRequestChannel())
	go readConnections(secondClient.GetWSocket(), secondClient.GetRequestChannel())
}

func readConnections(conn *websocket.Conn, requestChannel chan dtypes.Event) {
	for {
		event := dtypes.Event{}
		err := conn.ReadJSON(&event)
		log.Println("read json from readConnections", event.GetStr())
		if err != nil {
			log.Println("Error reading json.", err)
		}
		requestChannel <- event
	}
}

func serverReceiveComputations(firstClientRequestChannel, secondClientRequestChannel, firstRespondChannelServer, secondRespondChannelServer, requestChannelServer chan dtypes.Event) {
	// var latestState dtypes.Event
	for {
		log.Println("running servercomputations loop begining")
		select {
		case latestState := <-firstClientRequestChannel:
			log.Println("first client request chanel passed the info to server")
			requestChannelServer <- latestState
		case latestState := <-secondClientRequestChannel:
			log.Println("second client request channel passed the info to server")
			requestChannelServer <- latestState
		}
	}
}

func serverComputations(firstClientRequestChannel, secondClientRequestChannel, firstRespondChannelServer, secondRespondChannelServer, requestChannelServer chan dtypes.Event) {
	// var latestState dtypes.Event
	for {
		log.Println("Server received a event msg to compute at Server computations")
		playerPositions := <-requestChannelServer
		updatedPlayerPositions := handler.Handle(playerPositions)
		log.Println("updated positions of players ", updatedPlayerPositions.GetStr())
		updatedBotPositions := dijkstra.UpdateBots(updatedPlayerPositions)
		log.Println("upadted positions of bots", updatedBotPositions.GetStr())
		firstRespondChannelServer <- updatedBotPositions
		secondRespondChannelServer <- updatedBotPositions
		log.Println("Updated bot positons send to client objects")
	}
}
