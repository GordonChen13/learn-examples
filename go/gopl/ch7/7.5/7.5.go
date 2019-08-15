package main

import (
	"fmt"
	"io"
	"strings"
)

type IOLimitReader struct {
	r io.Reader
	n int64
}

func (r *IOLimitReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if n > int(r.n) {
		n = int(r.n)
	}
	err = io.EOF
	return
}

func LimitReader(r io.Reader, n int64) *IOLimitReader {
	var lr IOLimitReader
	lr.r = r
	lr.n = n
	return &lr
}

func main() {
	r := LimitReader(strings.NewReader("<html><body><h1>hello</h1></body></html>"), 21)
	buffer := make([]byte, 1024)
	n, err := r.Read(buffer)
	buffer = buffer[:n]
	fmt.Println(n, err, string(buffer))
}
