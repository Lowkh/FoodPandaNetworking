package main

import (
	"io"
	"net/http"
)

func GoMain(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go Main")
}

func GoTrack(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go Track")
}

func GoJS(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go JS")
}

func main() {
	/*
		http.Handle("/", c)
		http.Handle("/JS", a)
		http.Handle("/track", b)
	*/
	http.Handle("/", http.HandlerFunc(GoMain))
	http.Handle("/JS", http.HandlerFunc(GoJS))
	http.Handle("/track", http.HandlerFunc(GoTrack))

	http.ListenAndServe(":8080", nil)
}
