package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main () {
	start := time.Now()

	ch := make(chan string)

	args := os.Args[1:]

	for _, url := range args {
		go fetch(url, ch)
	}

	for range args {
		fmt.Println(<- ch)
	}

	seconds := time.Since(start).Seconds()

	fmt.Printf("all cost %.2f", seconds)
}

func fetch(url string, ch chan<- string) {

	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, response.Body)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	seconds := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs  %7d %s", seconds, nbytes, url)

	return
}
