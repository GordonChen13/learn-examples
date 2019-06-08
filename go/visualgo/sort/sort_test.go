package sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var cases []struct{
	name string
	input []int
	want []int
}

var int10 []int
var int100 []int
var int1000 []int
var int10000 []int
var int100000 []int

func init() {
	cases = [] struct{
		name string
		input []int
		want []int
	}{
		{"small int", []int{1, 8, 10, 7, 3, 88, 56, 93}, []int{1, 3, 7, 8, 10, 56, 88, 93}},
		{"big int", []int{10000, 8, 10, 7000, 3, 88, 5600000, 93}, []int{3, 8, 10, 88, 93, 7000, 10000, 5600000}},
	}

	int10 = intGen(10)
	int100 = intGen(100)
	int1000 = intGen(1000)
	int10000 = intGen(10000)
	int100000 = intGen(100000)
}

func intGen(num int) []int {
	res := make([]int, num)
	for i := 0; i < num; i++ {
		res = append(res, rand.Int())
	}

	return res
}

func TestBubble(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := Bubble(c.input)
			assert.Equal(t, c.want, get)
		})
	}
}

func BenchmarkBubble10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bubble(int10)
	}
}

func BenchmarkBubble100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bubble(int100)
	}
}

func BenchmarkBubble1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bubble(int1000)
	}
}

func BenchmarkBubble10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bubble(int10000)
	}
}

func BenchmarkBubble100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bubble(int100000)
	}
}


func TestSelect(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := Select(c.input)
			assert.Equal(t, c.want, get)
		})
	}
}

func BenchmarkSelect10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Select(int10)
	}
}

func BenchmarkSelect100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Select(int100)
	}
}

func BenchmarkSelect1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Select(int1000)
	}
}

func BenchmarkSelect10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Select(int10000)
	}
}

func BenchmarkSelect100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Select(int100000)
	}
}

func TestInsert(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := Insert(c.input)
			assert.Equal(t, c.want, get)
		})
	}
}

func BenchmarkInsert10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Insert(int10)
	}
}

func BenchmarkInsert100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Insert(int100)
	}
}

func BenchmarkInsert1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Insert(int1000)
	}
}

func BenchmarkInsert10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Insert(int10000)
	}
}

func BenchmarkInsert100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Insert(int100000)
	}
}

func TestMerge(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := MergeSort(c.input)
			assert.Equal(t, c.want, get)
		})
	}
}

func BenchmarkMerge10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(int10)
	}
}

func BenchmarkMerge100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(int100)
	}
}

func BenchmarkMerge1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(int1000)
	}
}

func BenchmarkMerge10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(int10000)
	}
}

func BenchmarkMerge100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(int100000)
	}
}
