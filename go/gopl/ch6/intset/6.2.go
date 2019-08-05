package intset

func (set *IntSet) AddAll(nums ...int)  {
	if len(nums) == 0 {
		return
	}
	for _, num := range nums {
		set.Add(num)
	}
	return
}
