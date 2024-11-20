package main

import "fmt"

func main() {
	name := "Welcome to Go Development"
	templateBasic := `
	<!DOCTYPE html>
	<html lang = "en'>
	<head>
	<meta charset = "UTF-8">
	<title>Hello!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(templateBasic)

}
