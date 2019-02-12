package play_node_runner

import (
	"dtypes"
	"fmt"

	handler "handlers"

	"github.com/gorilla/websocket"
)

// PlayNodeRunner is the event loop of NodeRunner
func PlayNodeRunner(conn *websocket.Conn) {
	for {
		event := dtypes.Event{}
		err := conn.ReadJSON(&event)

		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		handler.Handle(event)

		err = conn.WriteJSON(event)
		if err != nil {
			fmt.Print(err)
		}
	}
}
