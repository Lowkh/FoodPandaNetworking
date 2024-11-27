package main

import (
	"fmt"
	"os"

	"golang.org/x/net/websocket"
)

type msg struct {
	ID      int
	Message []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Expected Usage : ws://host:port")
	}
	service := os.Args[1]
	conn, err := websocket.Dial(service, "", "http://localhost")
	if err != nil {
		fmt.Println(err.Error())
	}

	message := msg{
		ID:      123,
		Message: []string{"First Msg", "Second Data", "Third Code"},
	}

	err = websocket.JSON.Send(conn, message)
	if err != nil {
		fmt.Println("Cannot Send!")
	}
}
