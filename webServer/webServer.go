// Demo: A simple webserver

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)           // each request calls handler
	http.HandleFunc("/count", countRequest) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

// handler echos requested URL path
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Counts the number of URL requests and reports
func countRequest(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "Count = %d\n", count)
}
