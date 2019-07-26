package main

import "fmt"

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1,2))
	fmt.Println(sum(1,2,3))
}

func sum(nums ...int) int {
	var total = 0
	for _, num := range nums {
		total+=num
	}

	return total
}
