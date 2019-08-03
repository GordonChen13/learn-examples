package intset

func (s *IntSet) Len() int {
	var l int
	if len(s.words) == 0 {
		l = 0
	}
	for _, word := range s.words{
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word & (1 << uint(j)) != 0 {
				l++
			}
		}
	}
	return l
}

func (s *IntSet) Remove(x int)  {
	if !s.Has(x) {
		return
	}
	i, word := x/64, x%64
	filter := ^(1 << uint(word))
	s.words[i] = s.words[i] & uint64(filter)
	return
}

func (s *IntSet) Clear()  {
	s.words = []uint64{}
	return
}

func (s *IntSet) Copy() *IntSet {
	set := &IntSet{}
	set = s
	return set
}