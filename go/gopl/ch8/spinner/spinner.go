package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45

	start2 := time.Now()
	fibN := fib2(n)
	fmt.Printf("\rfib2: Fibonacci(%d) = %d; duration: %s \n", n, fibN, time.Since(start2))

	start1 := time.Now()
	fibN = fib(n)
	fmt.Printf("\rfib1: Fibonacci(%d) = %d; duration: %s \n", n, fibN, time.Since(start1))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func fib2(x int) int {
	var nums = make([]int, x+1)
	for i := 0; i <= 45; i++ {
		if i < 2 {
			nums[i] = i
		} else {
			nums[i] = nums[i-1] + nums[i-2]
		}
	}
	return nums[x]
}
