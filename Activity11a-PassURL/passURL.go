package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", URLCall)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func URLCall(res http.ResponseWriter, req *http.Request) {
	passValue := req.FormValue("query")
	fmt.Fprintln(res, "Value is: "+passValue)
}

// http://localhost:8080/?query=anything
