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

func TestAdder(t *testing.T) {
	pos, neg := Adder(), Adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func TestPic(t *testing.T) {
	//pic.Show(Pic)
}

func TestWordCount(t *testing.T) {
	cases := []struct {
		in string
		want map[string]int
	}{
		{
			"I am learning Go!",
			map[string]int{"am":1, "learning":1, "Go!":1, "I":1},
		},
		{
			"I ate a donut. Then I ate another donut.",
			map[string]int{"Then":1, "another":1, "I":2, "ate":2, "a":1, "donut.":2},
		},
	}
	for _, c := range cases {
		got := WordCount(c.in)
		t.Logf("wordcount(%s) == %v, want %v", c.in ,got ,c.want)
	}
}

func TestFibonacciClo(t *testing.T) {
	f := FibonacciClo()
	for i := 0; i< 10 ; i ++ {
		num := f()
		fmt.Printf("%d ", num)
	}
}