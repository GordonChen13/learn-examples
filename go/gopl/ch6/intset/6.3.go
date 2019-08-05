package intset

func (s *IntSet) IntersectWith(t *IntSet)  {
	for i := range t.words {
		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet)  {
	for i := range t.words {
		s.words[i] &= ^t.words[i]
	}
}

func (s *IntSet) SymmetricWith(t *IntSet)  {
	for i := range t.words {
		s.words[i] ^= t.words[i]
	}
}