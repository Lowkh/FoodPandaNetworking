package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : ws://host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	header := make(http.Header)
	header.Add("Origin", "http://localhost:4321")
	conn, _, err := websocket.DefaultDialer.Dial(service, header)
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, reply, err := conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF from Server")
				break
			}

			if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				fmt.Println("Close from Server")
				break
			}
			fmt.Println("Cannot receive msg ", err.Error())
			break
		}
		fmt.Println("Received from Server : ", string(reply[:]))

		err = conn.WriteMessage(websocket.TextMessage, reply)
		if err != nil {
			fmt.Println("Cannot return msg")
		}

	}
}
