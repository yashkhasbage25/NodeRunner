package play_node_runner

import (
	"coord"
	"dijkstra"
	"dtypes"
	"fmt"
	handler "handlers"

	"github.com/gorilla/websocket"
)

// PlayNodeRunner is the event loop of NodeRunner
func PlayNodeRunner(conn *websocket.Conn) {
	coord.initialize()
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
