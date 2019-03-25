package main

import (
	"io"
	"net/http"
)

func main()  {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "Hello Gophercon!")
}
