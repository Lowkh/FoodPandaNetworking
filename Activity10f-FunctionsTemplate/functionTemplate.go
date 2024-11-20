package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type student struct {
	Name  string
	Class string
}

type totalScore struct {
	Module1 int
	Module2 int
}

var funcMap = template.FuncMap{
	"studentName":  getStudentName,
	"studentClass": getStudentClass,
	"score":        calScore,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("template.gohtml"))
}

func getStudentName(s student) string {
	return s.Name
}

func getStudentClass(s student) string {
	return s.Class
}

func calScore(ts totalScore) int {
	score := ts.Module1 + ts.Module2
	return score
}

func main() {

	student1 := student{
		Name:  "Obiwan",
		Class: "Tatooine",
	}

	student2 := student{
		Name:  "Lord Vader",
		Class: "Death Star",
	}

	score1 := totalScore{
		Module1: 10,
		Module2: 20,
	}
	score2 := totalScore{
		Module1: 30,
		Module2: 40,
	}
	students := []student{student1, student2}
	scores := []totalScore{score1, score2}

	data := struct {
		StudentData []student
		ScoreData   []totalScore
	}{
		students,
		scores,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
