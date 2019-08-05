package intset

func (s *IntSet) Elems() []int {
    var elems []int
    for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word & (1 << uint(j)) != 0 {
				elems = append(elems, 64*i + j)
			}
		}
	}

    return elems
}
