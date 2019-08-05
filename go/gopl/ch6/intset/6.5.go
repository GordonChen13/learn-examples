package intset

import (
	"bytes"
	"fmt"
)

type ISet2 struct {
	words []uint
}

const Bit = 32 << (^uint(0) >> 63)

func (s *ISet2) Has(x int) bool {
	word, bit := x/Bit, uint(x%Bit)
	return word < len(s.words) && s.words[word] & (1 << bit) != 0
}

func (s *ISet2) Add(x int)  {
	word, bit := x/Bit, uint(x%Bit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *ISet2) UnionWith(t *ISet2) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *ISet2) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < Bit; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", Bit*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
