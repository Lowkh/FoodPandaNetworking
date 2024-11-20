package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", form)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func form(res http.ResponseWriter, req *http.Request) {

	var stringToPrint string

	//This checks for if the method is POST.
	if req.Method == http.MethodPost {
		file, fileInfo, err := req.FormFile("filename") // Reads the filename and return the filename and file header multipart
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fmt.Println("\nFile:", file, "\nFile-HeaderProperties:", fileInfo, "\nerr: ", err)

		readData, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		stringToPrint = string(readData)

		destFile, err := os.Create(fileInfo.Filename)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer destFile.Close()

		_, err = destFile.Write(readData)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="filename">
	<input type="submit">
	</form>
	<br>`)
}
