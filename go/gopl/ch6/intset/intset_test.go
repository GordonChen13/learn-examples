package intset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSet_Has(t *testing.T) {
	set1 := &IntSet{}
	set1.Add(10)
	assert.Equal(t, true, set1.Has(10), "set1 should has 10")
	assert.Equal(t, false, set1.Has(20), "set1 should'nt has 20")

	set2 := &IntSet{}
	set2.Add(100)
	set1.UnionWith(set2)
	assert.Equal(t, true, set1.Has(100), "set1 should has 100")
}

func TestIntSet_String(t *testing.T) {
	set1 := &IntSet{}
	set1.Add(1)
	set1.Add(80)
	set1.Add(300)
	assert.Equal(t, "{1 80 300}", set1.String())

	set2 := &IntSet{}
	set2.Add(150)
	assert.Equal(t, "{150}", set2.String())
}

//BenchmarkIntSet_String-8         3000000               418 ns/op
func BenchmarkIntSet_String(b *testing.B) {
	set1 := &IntSet{}
	set1.Add(10)
	set1.Add(800)
	set1.Add(3000)
	for i := 0; i < b.N; i++ {
		set1.String()
	}
}

//BenchmarkIntSet_StringV2-8       3000000               549 ns/op
func BenchmarkIntSet_StringV2(b *testing.B) {
	set1 := &IntSet{}
	set1.Add(10)
	set1.Add(800)
	set1.Add(3000)
	for i := 0; i < b.N; i++ {
		set1.StringV2()
	}
}
