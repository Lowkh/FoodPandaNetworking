package main

import (
	"fmt"
	"log"
	"net"
)

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

func handle(myListener net.PacketConn, addr net.Addr, buffer []byte) {
	fmt.Println(string(buffer))
	myListener.WriteTo(buffer, addr)

}
