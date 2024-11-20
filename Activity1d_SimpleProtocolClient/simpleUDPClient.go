package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	startUDPClient()
}

func sendAlive(conn net.Conn) {
	for {
		time.Sleep(3 * time.Second)
		fmt.Fprintf(conn, "alive")
	}

}

func startUDPClient() {
	conn, err := net.Dial("udp", "localhost:5331")
	if err != nil {
		log.Fatalln("Connection fails")
		return
	}
	defer conn.Close()
	go sendAlive(conn)
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Key in your message: ")
		message, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, message+"\n")

		recMessage, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Received : ", recMessage)

	}
}
