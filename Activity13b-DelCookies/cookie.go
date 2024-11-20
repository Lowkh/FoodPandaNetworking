package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/delete", deleteCookie)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, `<h1><a href="/set">Create My Cookie</a></h1>`)
}

func setCookie(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "GoTrack",
		Value: "GS001",
	})

	http.SetCookie(res, &http.Cookie{
		Name:  "GoJS",
		Value: "second",
	})

	http.SetCookie(res, &http.Cookie{
		Name:  "Panda",
		Value: "Panda2",
	})

	fmt.Fprintln(res, `<h1><a href="/read">Read My Cookie</a></h1>`)
}

func readCookie(res http.ResponseWriter, req *http.Request) {
	myCookie, err := req.Cookie("GoTrack")
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(res, `<h1>Your Cookie:<br>%v</h1><h1><a href="/delete">Remove My Cookie</a></h1>`, myCookie)
}

func deleteCookie(res http.ResponseWriter, req *http.Request) {
	myCookie, err := req.Cookie("GoSchool")
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}
	myCookie.MaxAge = -1 // delete cookie
	http.SetCookie(res, myCookie)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
