package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func setCookie(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "GoTrackPanda",
		Value: "GST001",
		Path:  "/",
	})

	fmt.Fprintln(res, "To see the cookie, go to: dev tools / application / cookies")
}

func readCookie(res http.ResponseWriter, req *http.Request) {

	myCookie, err := req.Cookie("GoTrackPanda")
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(res, "your cookie content is ", myCookie)
}
