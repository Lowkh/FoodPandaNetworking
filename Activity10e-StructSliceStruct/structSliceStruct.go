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

type course struct {
	CourseName string
	MaxCount   int
}

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {

	student1 := student{
		FirstName: "Anakin",
		LastName:  "Skywalker",
	}

	student2 := student{
		FirstName: "R2",
		LastName:  "D2",
	}

	course1 := course{
		CourseName: "StarFighter Navigation",
		MaxCount:   50,
	}

	course2 := course{
		CourseName: "Deathstar Assault",
		MaxCount:   100,
	}

	students := []student{student1, student2}
	courses := []course{course1, course2}

	data := struct {
		StudentNames []student
		CourseNames  []course
	}{
		students,
		courses,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
