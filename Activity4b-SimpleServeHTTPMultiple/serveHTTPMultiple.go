package main

import (
	"fmt"
	"io"
	"net/http"
)

type GoTrack string
type GoJS int
type GoMain int

func (b GoTrack) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go Track")
	fmt.Println("Track")
}

func (a GoJS) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go JS")
	fmt.Println("Go JS")
}

func (c GoMain) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go Main")
	fmt.Println("Main")
}

func main() {
	var a GoTrack
	var b GoJS
	var c GoMain

	http.Handle("/", c)
	http.Handle("/JS", b)
	http.Handle("/Track", a)

	http.ListenAndServe(":8080", nil)
}
