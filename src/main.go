package main

import (
    "github.com/gorilla/websocket"
    "net/http"
    // "time"
    "io/ioutil"
    // "log"
    "fmt"
)

type Event struct {
    Etype string `json:"etype"`
    Px uint32 `json:"px"`
    Py uint32 `json:"py"`
}

type UpdatedPositions struct {
    Px uint32 `json:"px"`
    Py uint32 `json:"py"`

    // b1x uint32 `json:"b1x"`
    // b1y uint32 `json:"b1y"`
    // b2x uint32 `json:"b2x"`
    // b2y uint32 `json:"b2y"`
    //
    // g1x uint32 `json:"g1x"`
    // g1y uint32 `json:"g1y"`
    // g2x uint32 `json:"g2x"`
    // g2y uint32 `json:"g2y"`
}

func unlimited_write(conn *websocket.Conn) {
    for {
        // time.Sleep(1000 * time.Millisecond)
        // pos := UpdatedPositions{}
        event := Event{}
        // var msg int
        err := conn.ReadJSON(&event)

        if err != nil {
            fmt.Println("Error reading json.", err)
        }

        // fmt.Printf("Got message: %#v\n", event.P)
        fmt.Println(event.Px, event.Py, event.Etype)

        pos := UpdatedPositions{}
        if event.Etype == "up" {
            pos.Px = event.Px
            pos.Py = event.Py - 10
        } else if event.Etype == "down" {
            pos.Px = event.Px
            pos.Py = event.Py + 10
        }

        err = conn.WriteJSON(pos)
        if err != nil {
            fmt.Print(err)
        }
        // var rec string
        // err = conn.ReadJSON(&rec)
        // if err!=nil {
        //     log.Fatal(err)
        // } else {
        //     fmt.Println(rec)
        // }
    }
}

func main() {
    // upgrader := websocket.Upgrader{
	// 		ReadBufferSize: 2048,
	// 		WriteBufferSize: 2048,
	// 		CheckOrigin: func(r *http.Request) bool {
	// 			return true
	// 		}}
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        content, err := ioutil.ReadFile("index.html")
    	if err != nil {
    		fmt.Println("Could not open file.", err)
    	}
    	fmt.Fprintf(w, "%s", content)
    })
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
        if err != nil {
            fmt.Print(err)
        }
        fmt.Print("connecttion found")

        go unlimited_write(conn)
    })
    panic(http.ListenAndServe(":8080", nil))
}
