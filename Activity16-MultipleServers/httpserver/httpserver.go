package httpserver

import (
	"fmt"
	"net/http"
)

// StartHTTPServer is the start of HTTP Servers
func StartHTTPServer() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/about", http.HandlerFunc(about))

	http.ListenAndServe(":5331", nil)
}

func index(ress http.ResponseWriter, req *http.Request) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>INDEX</strong><br>
	<br><a href="/">Index</a><br>
	<a href="/about">About</a><br>
	</body></html>`
	ress.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(ress, body)
}

func about(ress http.ResponseWriter, req *http.Request) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>ABOUT</strong><br>
	<br>Welcome to Shoppee Challenge!<br>
	<br><a href="/">Index</a><br>
	</body></html>`
	ress.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(ress, body)
}
