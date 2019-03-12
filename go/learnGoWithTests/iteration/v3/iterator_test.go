package v1

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T)  {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1000)
	}
}

func BenchmarkPkgRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PkgRepeat("a", 1000)
	}
}
// 内存分配导致性能差异
//BenchmarkRepeat-8          10000            103905 ns/op
//BenchmarkPkgRepeat-8     5000000               357 ns/op

func ExampleRepeat() {
	bbb := Repeat("b", 3)
	fmt.Println(bbb)
	// Output: bbb
}
