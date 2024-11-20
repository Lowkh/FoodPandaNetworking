package customclient

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

//StartTCPClient starts a tcp client and waits for user input
func StartTCPClient() {
	conn, err := net.Dial("tcp", "localhost:5331")
	if err != nil {
		log.Fatalln("Connection fails")
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Key in your message: ")
		message, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, message+"\n")

		recMessage, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Received : ", recMessage)

	}
}
