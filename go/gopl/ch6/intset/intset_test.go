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
	set1.Add(2)
	set1.Add(3)
	assert.Equal(t, "{1 2 3}", set1.String())

	set2 := &IntSet{}
	set2.Add(999999999)
	assert.Equal(t, "{999999999}", set2.String())
}
