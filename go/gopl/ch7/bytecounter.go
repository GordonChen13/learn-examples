package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int ,error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)

	c = 0
	fmt.Fprintln(&c, "hello, world")
	fmt.Println(c)
}
