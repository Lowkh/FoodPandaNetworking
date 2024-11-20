package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template.gohtml"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"dbl":  double,
	"sq":   square,
	"sqrt": sqRoot,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", 5)
	if err != nil {
		log.Fatalln(err)
	}
}
