package tour

import (
	"testing"
	"fmt"
	"math/rand"
)

func TestPackages(t *testing.T) {
	Packages()
}

func BenchmarkPackages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Packages()
	}
}

func TestImport(t *testing.T) {
	Import()
}

func TestAdd(t *testing.T) {
	a := rand.Int()
	b := rand.Int()
	c := Add(a, b)
	fmt.Printf("%d + %d = %d", a, b, c)
}

func TestForSum(t *testing.T) {
	x := 10000
	ForSum(x)
}

func TestLoopSqrt(t *testing.T) {
	x := 10.0
	y := 100
	LoopSqrt(x, y)

}
