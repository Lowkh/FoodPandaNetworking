package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Welcome to Go Development"

	str := fmt.Sprint(`
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
	`)

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer newFile.Close()

	io.Copy(newFile, strings.NewReader(str))
}
