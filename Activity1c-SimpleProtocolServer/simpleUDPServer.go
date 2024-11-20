package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var aliveCount int = 10

func main() {
	startUDPServer()
}

func startUDPServer() {
	addr, err := net.ResolveUDPAddr("udp", ":5331")
	myListener, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer myListener.Close()
	go aliveCheck()
	for {
		buffer := make([]byte, 1024)

		len, addr, err := myListener.ReadFrom(buffer)

		if err != nil {
			fmt.Println("Continue")
			continue
		}

		go handle(myListener, addr, buffer[:len])
	}
}

func aliveCheck() {
	for {

		time.Sleep(500 * time.Millisecond)
		if aliveCount < 1 {
			fmt.Println("No client connected! Waiting")
			aliveCount = 0
		} else {
			aliveCount -= 1
		}
	}
}
func handle(myListener net.PacketConn, addr net.Addr, buffer []byte) {

	message := string(buffer)
	if message == "alive" {
		aliveCount = 10
		fmt.Println("Client alive message received")
	} else {
		fmt.Println(string(buffer))
		myListener.WriteTo(buffer, addr)
	}

}
