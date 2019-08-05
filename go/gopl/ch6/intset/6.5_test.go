package intset_test

import (
	"github.com/GordonChen13/learn-examples/go/gopl/ch6/intset"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestISet2_Has(t *testing.T) {
	set1 := &intset.ISet2{}
	set1.Add(10)
	assert.Equal(t, true, set1.Has(10), "set1 should has 10")
	assert.Equal(t, false, set1.Has(20), "set1 should'nt has 20")

	set2 := &intset.ISet2{}
	set2.Add(100)
	set1.UnionWith(set2)
	assert.Equal(t, true, set1.Has(100), "set1 should has 100")
}

func TestISet2_String(t *testing.T) {
	set1 := &intset.ISet2{}
	set1.Add(1)
	set1.Add(80)
	set1.Add(300)
	assert.Equal(t, "{1 80 300}", set1.String())

	set2 := &intset.ISet2{}
	set2.Add(150)
	assert.Equal(t, "{150}", set2.String())
}

