package main

import "fmt"

func main() {
	defer func() {
		 if num := recover(); num != nil {
		 	fmt.Println(num)
		 }
	}()
	noReturn(5)
}

func noReturn(num int) {
	panic(num)
}
