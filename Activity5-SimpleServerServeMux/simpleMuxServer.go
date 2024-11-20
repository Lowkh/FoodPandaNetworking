package main

import (
	"io"
	"net/http"
)

type GoTrack int
type GoJS int
type GoMain int

func (b GoTrack) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go Track")
}

func (a GoJS) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go JS")
}

func (c GoMain) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go Main")
}

func main() {
	var a GoJS
	var b GoTrack
	var c GoMain

	mux := http.NewServeMux()
	mux.Handle("/", c)
	mux.Handle("/JS", a)
	mux.Handle("/track", b)

	http.ListenAndServe(":8080", mux)
}
