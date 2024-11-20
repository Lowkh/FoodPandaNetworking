package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templateVariables.gohtml"))
}

func main() {

	modules := map[string]string{
		"1": "Go Basic",
		"2": "Go Advanced",
		"3": "Go In Action",
		"4": "Go Microservices",
	}

	err := tpl.Execute(os.Stdout, modules)
	if err != nil {
		log.Fatalln(err)
	}
}
