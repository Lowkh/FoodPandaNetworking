package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage ws://host:port")
		os.Exit(1)
	}
	service := os.Args[1]
	conn, _ := websocket.Dial(service, "", "http://localhost:8080")

	var msg string
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				//graceful shutdown
				break
			}
			fmt.Println("Could not receive msg " + err.Error())
			break
		}
		fmt.Println("From Server: " + msg)
		err = websocket.Message.Send(conn, msg)
		if err != nil {
			fmt.Println("Cannot return msg")
			break
		}
	}
	os.Exit(0)
}
