package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		file, err := os.OpenFile("output.html", os.O_RDWR, os.ModeAppend)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open file err: %v\n", err)
		}
		_, err = io.Copy(file, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
	}
}
