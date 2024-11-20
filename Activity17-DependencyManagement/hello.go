package main

import (
	"fmt"
	"rsc.io/quote"
)

//Make references to
//https://blog.golang.org/using-go-modules
func main() {
	message := Hello()
	fmt.Println(message)
}

func Hello() string {
	//return "Hello, world."
	return quote.Hello()
}
