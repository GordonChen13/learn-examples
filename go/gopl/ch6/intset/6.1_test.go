package intset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSet_Len(t *testing.T) {
	set1 := &IntSet{}
	set1.Add(100)
	set2 := &IntSet{}
	set2.Add(100)
	set2.Add(1000)
	set3 := &IntSet{}
	set3.Add(1)
	set3.Add(1000)
	set3.Add(1000000)
	cases := []struct{
		name string
		set *IntSet
		want int
	} {
		{"return 0 when empty", &IntSet{}, 0},
		{"return 1", set1, 1},
		{"return 2", set2, 2},
		{"return 3", set3, 3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.set.Len()
			assert.Equal(t, c.want, got)
		})
	}
}

func TestIntSet_Remove(t *testing.T) {
	set3 := &IntSet{}
	set3.Add(1)
	set3.Add(1000)
	set3.Add(1000000)

	set3.Remove(1)
	assert.Equal(t, false, set3.Has(1))
	set3.Remove(1000)
	assert.Equal(t, false, set3.Has(1000))
	assert.Equal(t, true, set3.Has(1000000))
}

func TestIntSet_Clear(t *testing.T) {
	set3 := &IntSet{}
	set3.Add(1)
	set3.Add(1000)
	set3.Add(1000000)

	set3.Clear()
	assert.Equal(t, false, set3.Has(1))
	assert.Equal(t, false, set3.Has(1000))
	assert.Equal(t, false, set3.Has(1000000))
}

func TestIntSet_Copy(t *testing.T) {
	set3 := &IntSet{}
	set3.Add(1)
	set3.Add(1000)
	set3.Add(1000000)

	set4 := set3.Copy()
	assert.Equal(t,true, set4.Has(1))
	assert.Equal(t,true, set4.Has(1000))
	assert.Equal(t,true, set4.Has(1000000))
}
