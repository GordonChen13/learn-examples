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
	http.HandleFunc("/incr", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Paht = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, h *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "current count: %d", count)
	mu.Unlock()
}
