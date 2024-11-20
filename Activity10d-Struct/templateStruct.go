package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type student struct {
	FirstName string
	LastName  string
}

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {

	student1 := student{
		FirstName: "Obiwan",
		LastName:  "Kenobi",
	}

	err := tpl.Execute(os.Stdout, student1)
	if err != nil {
		log.Fatalln(err)
	}

}
