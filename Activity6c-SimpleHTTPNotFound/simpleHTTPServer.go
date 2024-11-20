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
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/feature1", feature1)
	http.HandleFunc("/feature2", feature2)
	http.ListenAndServe(":8080", nil)
}
