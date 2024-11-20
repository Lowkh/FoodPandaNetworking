package main

import (
	"fmt"
	"net/http"
	"time"
)

func sleepServer(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Sleeping.....")
	time.Sleep(time.Hour)
}

func main() {
	http.HandleFunc("/", sleepServer)
	http.ListenAndServe(":5331", nil)
}
