package main

import (
	"fmt"
	"io"
	"os"
)

type ByteWriter struct {
	writer io.Writer
	count  int64
}

func (w *ByteWriter) Write(p []byte) (int, error) {
	count, err := w.writer.Write(p)
	w.count += int64(count)

	return count, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bw := ByteWriter{w, 0}
	return &bw, &bw.count
}

func main() {
	writer, count := CountingWriter(os.Stdout)
	fmt.Fprint(writer, "hello world\n")
	fmt.Println(*count)
}
