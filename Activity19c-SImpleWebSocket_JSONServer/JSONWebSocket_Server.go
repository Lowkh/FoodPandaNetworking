package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

type msg struct {
	ID      int
	Message []string
}

func ReceiveMessage(ws *websocket.Conn) {
	var message msg
	err := websocket.JSON.Receive(ws, &message)
	if err != nil {
		fmt.Println("Can't Receive!")
	} else {
		fmt.Println("Message ID: ", message.ID)
		for _, content := range message.Message {
			fmt.Println("Message Content: ", content)
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(ReceiveMessage))
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}
