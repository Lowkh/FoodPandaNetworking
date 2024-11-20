package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
	UserID   string
}

var tpl *template.Template
var mapUsers = map[string]user{}
var mapSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/restricted", restricted)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		id := uuid.NewV4()
		myCookie = &http.Cookie{
			Name:  "myCookie",
			Value: id.String(),
		}
		http.SetCookie(res, myCookie)
	}

	var myUser user
	if username, ok := mapSessions[myCookie.Value]; ok {
		myUser = mapUsers[username]
	}

	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		myUser = user{username, firstname, lastname, myCookie.Value}
		mapSessions[myCookie.Value] = username
		mapUsers[username] = myUser
	}

	tpl.ExecuteTemplate(res, "index.gohtml", myUser)
}

func restricted(res http.ResponseWriter, req *http.Request) {

	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	username, ok := mapSessions[myCookie.Value]
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	myUser := mapUsers[username]

	if myUser.First != "" {
		tpl.ExecuteTemplate(res, "restricted.gohtml", myUser)
	} else {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

}
