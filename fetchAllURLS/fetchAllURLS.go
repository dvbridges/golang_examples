// Demo: Fetch multiple URLs using concurrency

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a go routine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2f time elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // Dont leak resources
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("URL: %v\nTime: %.2f seconds\nSize: %d bytes\n", url, secs, nBytes)

}
