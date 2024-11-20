package main

import (
	httpServerCustom "Activity16-MultipleServers/httpserver"
	"fmt"
)

func main() {
	fmt.Println("TCP server")
	//	customclient.StartTCPClient()
	//jsonServerCustom.StartJSONQuery()
	httpServerCustom.StartHTTPServer()
}
