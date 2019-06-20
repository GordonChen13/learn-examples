package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseBytes(a []byte) []byte {
	if utf8.RuneCount(a) == 1 {
		return a
	}
	_, s := utf8.DecodeRune(a)
	return append(reverseBytes(a[s:]), a[:s]...)
}

func main() {
	a := []byte("你好吗？abc?")
	fmt.Println(string(reverseBytes(a)))
}
