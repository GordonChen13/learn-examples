package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var t = flag.String("e", "sha256", "encrypt method, only support sha256,sha384 and sha512")
	flag.Parse()

	input := []byte("x")
	fmt.Printf("%s\n",string(input[:]))
	var res string
	switch *t {
	case "sha256":
		o := sha256.Sum256(input)
		res = string(o[:])
	case "sha384":
		o := sha512.Sum384(input)
		res = string(o[:])
	case "sha512":
		o := sha512.Sum512(input)
		res = string(o[:])
	default:
		o := sha256.Sum256(input)
		res = string(o[:])
	}

	fmt.Printf("%x\n",res)
}
