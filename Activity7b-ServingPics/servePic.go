package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/germany", kyoto)
	http.HandleFunc("/kyotoOnline", kyotoOnline)
	http.ListenAndServe(":8080", nil)
}

func kyoto(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="germany.jpg">`)
}

func kyotoOnline(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/2/22/Kyoto_montage.jpg">`)
}
