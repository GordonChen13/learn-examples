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

func LoopSqrt(x float64, y int) float64{
	var z = 1.0
	for i :=0; i < y; i ++ {
		z -= (z*z - x) / (2*z)
	}
	fmt.Printf("the sqrt of %f is %f, runs %d times\n", x, z, y)
	return z;
}

func Adder() func(int) int  {
	sum := 5
	return func(x int) int {
		sum = sum + x
		return sum
	}
}