package intset

import (
	"bytes"
	"fmt"
	"strconv"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
   word, bit := x/64, uint64(x%64)
   return word < len(s.words) && s.words[word] & (1 << bit) != 0
}

func (s *IntSet) Add(x int)  {
	word, bit := x/64, uint64(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//BenchmarkIntSet_String-8         3000000               418 ns/op
func (s *IntSet) String() string {
	str := "{"
	for word, bit := range s.words {
		if bit == 0 {
			continue
		}
		if word != 0 && str != "{"{
			str += " "
		}
		for i := 0; i< 64; i++ {
			if bit & (1 << uint64(i)) != 0 {
				num := 64*word + i
				str += strconv.FormatInt(int64(num), 10)
			}
		}
	}
	str += "}"
	return str
}

//BenchmarkIntSet_StringV2-8       3000000               549 ns/op
func (s *IntSet) StringV2() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
