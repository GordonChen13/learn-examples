package popcount

import (
	"fmt"
	"testing"
)

func TestPopCount(t *testing.T) {
	count := PopCount(10000)
	fmt.Println(count)
}

func BenchmarkPopCount10(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount(10)
	}
}

func BenchmarkPopCount10000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount(10000)
	}
}

func BenchmarkPopCount100000000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount(100000000)
	}
}
