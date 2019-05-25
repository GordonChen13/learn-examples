package popcount

func PopCount3(x uint64) int {
	var count byte
	for i := 0; i < 8; i ++ {
		count += pc[byte(x>>(byte(i)*8))]
	}

	return int(count)
}
