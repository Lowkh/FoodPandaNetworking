package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	/*
		fmt.Println("Requesting information")
		data, err := http.Get("http://127.0.0.1:5331")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	*/

	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	_, err := c.Get("http://127.0.0.1:5331")
	if err != nil {
		fmt.Println(err)
	}
}
