package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	g1 := struct {
		Value1 int
		Value2 int
	}{
		7,
		9,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", g1)
	if err != nil {
		log.Fatalln(err)
	}
}
