package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup

	mutex sync.Mutex
)

func incCounter(id string) {
	defer wg.Done()
	for count := 0; count < 10; count++ {
		//mutex.Lock() //-- disable to see interleaving effect
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
			fmt.Println(id, ":", counter)
			fmt.Println()
		}
		//mutex.Unlock() //-- disable to see interleaving effect
	}
}
func feature1(res http.ResponseWriter, req *http.Request) {
	wg.Add(2)
	go incCounter("Feature 1")
	go incCounter("Feature 1.5")
	feature2(res, req) //-- enable to simulate effect of another client calling feature2
	wg.Wait()
	io.WriteString(res, strconv.Itoa(counter))
}
func feature2(res http.ResponseWriter, req *http.Request) {
	wg.Add(1)
	go incCounter("Feature 2")
	wg.Wait()
	io.WriteString(res, strconv.Itoa(counter))

}
func main() {
	http.HandleFunc("/feature1", feature1)
	http.HandleFunc("/feature2", feature2)
	http.ListenAndServe(":8080", nil)

}
