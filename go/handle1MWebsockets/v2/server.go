package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func ws(w http.ResponseWriter, r *http.Request)  {
	// Upgrade connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	for  {
		_, msg, err := conn.ReadMessage()
		if err != nil {
            log.Printf("Failed to read message %v", err)
            conn.Close()
			return
		}
		println(msg)
	}
}

func main()  {
	http.HandleFunc("/", ws)
	http.ListenAndServe(":8000", nil)
}
