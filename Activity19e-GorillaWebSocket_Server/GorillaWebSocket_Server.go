package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func HandleGorilla(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for n := 0; n < 10; n++ {
		msg := "Hello " + string(n+48)
		fmt.Println("Sending to Client: ", msg)
		err = conn.WriteMessage(websocket.TextMessage, []byte(msg))

		_, reply, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Can't Receive")
			break
		}
		fmt.Println("Received back from Client: " + string(reply[:]))
	}
	conn.Close()
}

func main() {
	http.HandleFunc("/", HandleGorilla)
	err := http.ListenAndServe(":4321", nil)
	if err != nil {
		fmt.Println(err)
	}
}
