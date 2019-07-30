package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	url := os.Args[1]
	if !strings.HasPrefix(url, "http") {
		log.Fatal("url must start with http")
	}
	name, n , err := fetch(url)
	fmt.Printf("name: %s, n: %d, err: %v\n", name, n, err)
}

func fetch(url string) (fileName string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("get url content error: %s", err)
	}
	defer resp.Body.Close()

	fileName = path.Base(resp.Request.URL.Path)
	if fileName == "/" || fileName == "." {
		fileName = "index.html"
	}
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("create file error: %s", err)
		return "", 0 , err
	}
	n, err = io.Copy(f, resp.Body)

	if closeErr := f.Close(); closeErr != nil {
		log.Fatalf("file close error: %s", closeErr)
        err = closeErr
	}
	return fileName, n, err
}
