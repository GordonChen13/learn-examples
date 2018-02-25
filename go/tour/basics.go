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