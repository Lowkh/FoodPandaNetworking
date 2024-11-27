package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error Upgrading Connection", err)
		return
	}
	defer conn.Close()

	fmt.Println("New Connection Established!")

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		fmt.Println("Received: %s\n", msg)

		if err = conn.WriteMessage(messageType, msg); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}

}

func main() {
	http.HandleFunc("/ws", handleConnections)
	fmt.Println("WebSocket server started!")
	http.ListenAndServe(":8080", nil)
}
