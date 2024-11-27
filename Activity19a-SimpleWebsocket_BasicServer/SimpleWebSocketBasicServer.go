package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", websocket.Handler(Echo))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Err")
	}

}

func Echo(ws *websocket.Conn) {
	fmt.Println("Echoing")
	for n := 0; n < 10; n++ {
		msg := "Hello " + string(n+48)
		fmt.Println("Sending to Client! " + msg)
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			fmt.Println("Can't send!")
			break
		}

		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Can't Receive")
			break
		}
		fmt.Println("Replies :" + reply)
	}
}
