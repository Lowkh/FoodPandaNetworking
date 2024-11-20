package main

import (
	"fmt"
	"io"
	"net/http"

	uuid "github.com/satori/go.uuid"
	//"github.com/gofrs/uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("myCookie")
	fmt.Println(err)
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:   "myCookie",
			Value:  id.String(),
			MaxAge: 1,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)
	io.WriteString(res, cookie.Name+"\n")
	io.WriteString(res, cookie.Value)
}
