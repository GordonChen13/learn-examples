package main

import (
	"fmt"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	fmt.Println("before sleep")
	time.Sleep(5*time.Second)
	fmt.Println("after sleep")
}

func trace(name string) func() {
	start := time.Now()
	fmt.Println(name + " star running")
	f := func() {
		d := time.Since(start)
		fmt.Printf("%s end running, time: %s\n", name, d)
	}
	return f
}
