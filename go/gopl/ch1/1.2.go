package main

import (
	"fmt"
	"os"
)

func main()  {
	for i, str := range os.Args[1:] {
		fmt.Println(i, str)
	}
}
