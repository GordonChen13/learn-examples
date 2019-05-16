package main

import (
"fmt"
"os"
)

func main()  {
	var s, sep string
	for _, str := range os.Args[1:] {
		s += sep + str
		sep = " "
	}

	fmt.Println(s)
}
