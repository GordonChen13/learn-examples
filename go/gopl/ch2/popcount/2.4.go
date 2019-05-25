package popcount

func PopCount4(x uint64) int {
	var count int
	for i := 0; i < 64; i ++ {
		count += int(x>>byte(i)&1)
	}

	return count
}
