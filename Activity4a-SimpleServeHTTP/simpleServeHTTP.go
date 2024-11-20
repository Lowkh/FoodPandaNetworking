package main

import (
	"fmt"
	"net/http"
)

// ServeHTTPPage tagged to a type int
type ServeHTTPPage int

func (m ServeHTTPPage) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("GoTrackKey", "This is from Go Track")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Your code in this func will run</h1>")
}

func main() {
	var d ServeHTTPPage
	http.ListenAndServe(":8080", d)
}
