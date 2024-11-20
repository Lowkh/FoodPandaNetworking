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

	modules := []string{"Go Basic", "Go Advanced", "Go In Action", "Go Microservices"}

	err := tpl.Execute(os.Stdout, modules)
	if err != nil {
		log.Fatalln(err)
	}
}
