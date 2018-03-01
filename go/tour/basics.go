package tour

import (
	"fmt"
	"math"
	"math/rand"
)

func Packages() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

func Import() {
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))
}

func Add(x int, y int) int {
	return x + y
}

func ForSum(x int) int {
	sum := 0
	fmt.Printf("the result of sum from 1 to %d is ", x)
	for i := 1; i <= x; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
	return sum
}