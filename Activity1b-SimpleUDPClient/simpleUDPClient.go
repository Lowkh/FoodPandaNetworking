package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	startUDPClient()
}

func startUDPClient() {
	conn, err := net.Dial("udp", "localhost:5331")
	if err != nil {
		log.Fatalln("Connection fails")
		return
	}
	defer conn.Close()

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Key in your message: ")
		message, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, message+"\n")

		recMessage, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Received : ", recMessage)

	}
}
