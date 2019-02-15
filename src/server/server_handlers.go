package server

import (
    "fmt"
	"io/ioutil"
    "net/http"
    "github.com/gorilla/websocket"
    play "play_node_runner"   

)

func (s *Server) SetHandlers() {
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
