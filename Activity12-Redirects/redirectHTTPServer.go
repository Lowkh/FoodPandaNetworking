package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/redirect", redirectFunction)
	http.HandleFunc("/request", requestFunction)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func requestFunction(res http.ResponseWriter, req *http.Request) {
	fmt.Println("requestFunction with:", req.Method, "Method")
	//http.Redirect(res, req, "/redirect", http.StatusSeeOther)
	//http.Redirect(res, req, "/redirect", http.StatusMovedPermanently)

}

func redirectFunction(res http.ResponseWriter, req *http.Request) {
	fmt.Println("redirected to here with:", req.Method, "Method")

}
