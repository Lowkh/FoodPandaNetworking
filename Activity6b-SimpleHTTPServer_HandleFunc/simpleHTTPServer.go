package main

import (
	"io"
	"net/http"
)

func feature1(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Feature1")
}

func feature2(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Feature2")
}
func main() {
	/*
		http.Handle("/", c)
		http.Handle("/JS", a)
		http.Handle("/track", b)

		http.Handle("/", http.HandlerFunc(GoMain))
		http.Handle("/JS", http.HandlerFunc(GoJS))
		http.Handle("/track", http.HandlerFunc(GoTrack))
	*/
	http.HandleFunc("/feature1", feature1)
	http.HandleFunc("/feature2", feature2)
	http.ListenAndServe(":8080", nil)
}
