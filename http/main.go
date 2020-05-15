// A simple http Listen and Serve demo from golang.org

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<h1>Hello World</h1>`)
}
