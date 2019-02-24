package server

import (
	"client"
	"dtypes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	play "play_node_runner"
	"strconv"

	"github.com/gorilla/websocket"
)

func SetHandlers(gameServer *Server) {
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

	http.HandleFunc("/web/assets/img/front.png", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/assets/img/front.png")
		if err != nil {
			fmt.Println("Could not open image.", err)
		}
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/css/index.css", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/css/index.css")
		if err != nil {
			fmt.Println("Could not open image.", err)
		}
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/css/wait.css", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/css/wait.css")
		if err != nil {
			fmt.Println("Could not open image.", err)
		}
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", content)
	})
}

// SetHandlers is sets all possible handlers for the server.
func (gameServer *Server) SetHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexContent, err := ioutil.ReadFile("web/index.html")
		if err != nil {
			fmt.Println("Could not open file.", err)
		}
		log.Println("Handling pattern /")
		fmt.Fprintf(w, "%s", indexContent)
	})

	http.HandleFunc("/web/wait.html", func(w http.ResponseWriter, r *http.Request) {
		waitContent, err := ioutil.ReadFile("web/wait.html")
		if err != nil {
			fmt.Println("Could not open file.", err)
		}
		fmt.Fprintf(w, "%s", waitContent)
		log.Println("handling pattern /web/wait.html")
	})

	http.HandleFunc("/wait", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			log.Fatal("Could not upgrade to websocket at web/wait.html (wait.html)", err)
		}
		log.Println("Websocket connection upgraded at wait.js")

		ip, port, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println("Could not separate ip and port.")
		}

		log.Println("server client count is", gameServer.GetNextID())
		if gameServer.CheckClientLimit() {
			newClient := &client.Client{
				IP:      ip,
				Port:    port,
				ID:      gameServer.GetNextID(),
				WSocket: conn,
			}
			log.Println("New client object created.", newClient.GetInfoStr())
			gameServer.AddNewClient(newClient)
			log.Println("New server state: ", gameServer.GetInfoStr())
		}
	})

	http.HandleFunc("/web/game.html", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/game.html")
		if err != nil {
			log.Println("Could not read file /web/game.html", err)
		}
		log.Println("handling pattern /web/game.html")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {

		ip, port, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println("Could not separate ip and port.")
		}

		conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			fmt.Print(err)
		}

		if gameServer.CheckClientLimit() {
			log.Println("Before getting sockets server is:", gameServer.GetInfoStr())
			thisClientID := gameServer.GetNextID()
			newClient := &client.Client{
				IP:      ip,
				Port:    port,
				ID:      thisClientID,
				WSocket: conn,
			}
			log.Println("New client object created.", newClient.GetInfoStr())
			gameServer.AddNewClient(newClient)
			log.Println("New state of server: ", gameServer.GetInfoStr())
			conn.WriteJSON(dtypes.Event{
				EventType: "SetClientID",
				Object:    strconv.Itoa(int(thisClientID)),
			})
			// gameServer.AddNewClient(newClient)
		}

		log.Println("handling pattern /game")
		// go play.PlayNodeRunner(conn)
	})

	http.HandleFunc("/web/assets/img/front.png", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/assets/img/front.png")
		if err != nil {
			log.Println("Could not open image.", err)
		}
		log.Println("handling pattern /web/assets/img/front.png")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/web/css/index.css", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/css/index.css")
		if err != nil {
			log.Println("Could not open image.", err)
		}
		w.Header().Add("Content-Type", "text/css")
		log.Println("handling pattern /web/css/index.css")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/web/css/wait.css", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/css/wait.css")
		if err != nil {
			log.Println("Could not open image.", err)
		}
		w.Header().Add("Content-Type", "text/css")
		log.Println("handling pattern /web/css/wait.css")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/web/css/game.css", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/css/game.css")
		if err != nil {
			log.Println("Could not open image.", err)
		}
		w.Header().Add("Content-Type", "text/css")
		log.Println("handling pattern /web/css/game.css")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/web/js/index.js", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/js/index.js")
		if err != nil {
			log.Println("Could not read file web/js/index.js")
		}
		w.Header().Add("Content-Type", "text/javascript")
		log.Println("handling pattern /web/js/index.js")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/web/js/wait.js", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/js/wait.js")
		if err != nil {
			log.Println("Could not read file web/js/wait.js")
		}
		w.Header().Add("Content-Type", "text/javascript")
		log.Println("handling pattern /web/js/wait.js")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/web/js/game.js", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/js/game.js")
		if err != nil {
			log.Println("Could not read file web/js/game.js")
		}
		w.Header().Add("Content-Type", "text/javascript")
		log.Println("handling pattern /web/js/wait.js")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/web/assets/svg/hourglass.svg", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/assets/svg/hourglass.svg")
		if err != nil {
			log.Println("Could not read file web/assets/svg/hourglass.svg")
		}
		w.Header().Add("Content-Type", "image/svg+xml")
		w.Header().Add("Vary", "Accept-Encoding")
		log.Println("handling pattern /web/assets/svg/hourglass.svg")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/web/assets/svg/level1.svg", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/assets/svg/level1.svg")
		if err != nil {
			log.Println("Could not read file web/assets/svg/level1.svg")
		}
		w.Header().Add("Content-Type", "image/svg+xml")
		w.Header().Add("Vary", "Accept-Encoding")
		log.Println("handling pattern /web/assets/svg/level1.svg")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/web/assets/img/player1.png", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/assets/img/player1.png")
		if err != nil {
			log.Println("Could not open image.", err)
		}
		log.Println("handling pattern /web/assets/img/player1.png")
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/web/assets/img/player2.png", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("web/assets/img/player2.png")
		if err != nil {
			log.Println("Could not open image.", err)
		}
		log.Println("handling pattern /web/assets/img/player2.png")
		fmt.Fprintf(w, "%s", content)
	})
}
