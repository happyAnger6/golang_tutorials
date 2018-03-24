package main

import (
	"net/http"
	"fmt"
	"log"
	"sync"
)

var mux sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	fmt.Fprintf(w, "total count:%d", count)
	mux.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	count++
	mux.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
