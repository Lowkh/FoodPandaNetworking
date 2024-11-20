package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", form)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func form(res http.ResponseWriter, req *http.Request) {
	fn := req.FormValue("firstName")
	ln := req.FormValue("lastName")
	g := req.FormValue("grade")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	//can try for using GET to see the different effect in url

	io.WriteString(res, `
	<form method="POST">
    <h1>Please enter the details</h1>
    <label for="firstName">First Name: </label>
    <input type="text" id="firstName" name="firstName">
    <br>
    <label for="lastName">Last Name: </label>
    <input type="text" id="lastName" name="lastName">
    <br>
    <label for="sub">Grade: </label>
    <input type="text" id="grade" name="grade">
    <br>
    <input type="submit">
	</form>
	<br>`)
	io.WriteString(res, `<br>`+fn)
	io.WriteString(res, `<br>`+ln)
	io.WriteString(res, `<br>`+g)
}
