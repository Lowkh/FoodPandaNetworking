package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5221", http.FileServer(http.Dir("."))))
}
