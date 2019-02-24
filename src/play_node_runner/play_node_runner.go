package play_node_runner

import (
	"coord"
	"dijkstra"
	"dtypes"
	"fmt"
	handler "handlers"
	"time"

	"github.com/gorilla/websocket"
)

func regularUpdater() {

	ticker := time.NewTicker(constants.UpdateRequestInterval) // Timer
	go func() {
		for range ticker.C {

		}
	}()
}

// PlayNodeRunner is the event loop of NodeRunner
func PlayNodeRunner(conn *websocket.Conn) {
	coord.initialize()
	health.SetHealths(1000, 1000)
	health.SetDecay(1, 500)
	go health.DecayHealth()
	go regularUpdater()
	for {
		event := dtypes.Event{}
		err := conn.ReadJSON(&event)

		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		updatedPlayerPositions := handler.Handle(event)
		updatedBotPositions := dijkstra.UpdateBots(updatedPlayerPositions)

		err = conn.WriteJSON(updatedBotPositions)
		if err != nil {
			fmt.Print(err)
		}
	}
}
