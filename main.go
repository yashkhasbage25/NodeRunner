package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	play "play_node_runner"

	"github.com/gorilla/websocket"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexContent, err := ioutil.ReadFile("web/index.html")
		if err != nil {
			fmt.Println("Could not open file.", err)
		}
		fmt.Fprintf(w, "%s", indexContent)
	})

	http.HandleFunc("/wait_to_join", func(w http.ResponseWriter, r *http.Request) {
		waitContent, err := ioutil.ReadFile("web/wait.html")
		if err != nil {
			fmt.Println("Could not open file.", err)
		}
		fmt.Fprintf(w, "%s", waitContent)
	})

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print("connecttion found")

		go play.PlayNodeRunner(conn)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
