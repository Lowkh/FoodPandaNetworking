package main

import (
	"io"
	"net/http"
)

// GoMenu defined as random type
type GoMenu int

func (m GoMenu) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		io.WriteString(res, "Welcome to Go Main")
	case "/JS":
		io.WriteString(res, "Welcome to Go JS")
	case "/track":
		io.WriteString(res, "Welcome to Go Track")
	}
}

func main() {

	var a GoMenu
	http.ListenAndServe(":8080", a)
}
