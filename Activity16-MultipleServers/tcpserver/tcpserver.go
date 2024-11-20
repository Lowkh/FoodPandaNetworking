package tcpserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// StartTCPServer is for starting of TCP Server Method
func StartTCPServer() {
	myListener, err := net.Listen("tcp", ":5331")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer myListener.Close()

	for {
		conn, err := myListener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received :", string(data))
		retMsg := string(data) + "\n"
		conn.Write([]byte(retMsg))
	}
}
