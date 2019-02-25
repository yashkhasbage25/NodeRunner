package play_node_runner

import (
	"github.com/IITH-SBJoshi/concurrency-3/src/client"
	"github.com/IITH-SBJoshi/concurrency-3/src/coords"
	"github.com/IITH-SBJoshi/concurrency-3/src/dijkstra"
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
	"fmt"
	handler "github.com/IITH-SBJoshi/concurrency-3/src/handlers"
	"time"
	"github.com/IITH-SBJoshi/concurrency-3/src/platform"

	"github.com/gorilla/websocket"
)

func regularUpdater(conn *websocket.Conn, requestChannelClient chan dtypes.Event) {

	ticker := time.NewTicker(50 * time.Millisecond)
	go func() {
		for range ticker.C {
			event := dtypes.Event{
				EventType: "SendUpdate",
			}

			err := conn.WriteJSON(event)
			if err != nil {
				fmt.Println("Error writing json.", err)
			}
			err := conn.ReadJSON(&event)
			if err != nil {
				fmt.Println("Error reading json.", err)
			}
			requestChannelClient <- event
		}
	}()
}

func sendResponse(receiveChannelClient chan dtypes.Event, conn *websocket.Conn) {
	responseMsg := <-receiveChannelClient
	conn.WriteJSON(responseMsg)
}

// PlayNodeRunner is the event loop of NodeRunner
func PlayNodeRunner(requestChannelServer, firstRespondChannelServer, secondRespondChannelServer chan dtypes.Event, firstClient, secondClient *client.Client) {
	coords.Initialize()
	platform.Initialize()
	health.SetHealths(1000, 1000)
	health.SetDecay(1, 500)
	go health.DecayHealth()
	go regularUpdater(firstClient.GetWSocket(), firstClient.GetRequestChannel())
	go regularUpdater(secondClient.GetWSocket(), secondClient.GetRequestChannel())
	go serverComputations(firstClient.GetRequestChannel(), secondClient.GetRequestChannel(), firstRespondChannelServer, secondRespondChannelServer, requestChannelServer)
	go sendResponse(firstClient.GetReceiveChannel(), firstClient.GetWSocket())
	go sendResponse(secondClient.GetReceiveChannel(), secondClient.GetWSocket())
	go readConnections(firstClient.GetWSocket(), firstClient.GetRequestChannel())
	go readConnections(secondClient.GetWSocket(), secondClient.GetRequestChannel())
	}
}

func readConnections(conn *websocket.Conn, requestChannel) {
	for {
		event := dtypes.Event{}
		err := conn.ReadJSON(&event)
		if err != nil {
			log.Println("Error reading json.", err)
		}
		requestChannel <- event
	}
}

func serverComputations(firstClientRequestChannel, secondClientRequestChannel, firstRespondChannelServer, secondRespondChannelServer, requestChannelServer chan dtypes.Event) {
	var latestState dtypes.Event
	for {
		select {
		case latestState <- firstClientRequestChannel:
			requestChannelServer <- latestState
		case latestState <- secondClientRequestChannel:
			requestChannelServer <- latestState
		}
		updatedPlayerPositions := handler.Handle(<-requestChannelServer)
		updatedBotPositions := dijkstra.UpdateBots(updatedPlayerPositions)
		firstRespondChannelServer <- updatedBotPositions
		secondRespondChannelServer <- updatedBotPositions
	}
}
