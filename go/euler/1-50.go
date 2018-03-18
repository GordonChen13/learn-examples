package euler

import (
	"fmt"
)

//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.
func MultiplesOf3And5()  {
	var sum int = 0;
	for i := 0; i < 1000; i ++ {
		if i % 5 == 0 || i % 3 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

//Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
func EvenFibonacciNumbers()  {
	sum := 0
	ch := make(chan int)
	go fibonacci(ch)
	for i := range ch {
		if i % 2 ==0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
func fibonacci(ch chan int)  {
	var x,y int = 1,2
	for {
		if x > 4000000 {
			close(ch)
			return
		}
		ch <- x
		x,y = y,x+y
	}
	close(ch)
}
