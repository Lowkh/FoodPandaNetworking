package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName string
	LastName  string
	Grade     string
}

func main() {
	http.HandleFunc("/", form)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func form(res http.ResponseWriter, req *http.Request) {

	firstname := req.FormValue("firstName")
	lastname := req.FormValue("lastName")
	grade := req.FormValue("grade")

	/* Add DB / DS query */

	err := tpl.ExecuteTemplate(res, "index.gohtml", person{firstname, lastname, grade})
	if err != nil {
		log.Fatalln(err)
	}
}
