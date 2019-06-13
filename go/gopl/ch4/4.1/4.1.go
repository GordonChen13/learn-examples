package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	var diff int
	for i := range c1 {
		diff += DiffBitCount(c1[i], c2[i])
	}

	fmt.Println(diff)
}

func DiffBitCount(b1, b2 byte) int{
	diff := int(b1^b2)
	return int(pc[diff])
}
