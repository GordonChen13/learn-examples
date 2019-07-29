package main


func Max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	// max 默认是0，需要第一个值覆盖
	var max int
	for i, n := range nums {
		if n >= max || i == 0 {
			max = n
		}
	}
	return max
}

func Min(nums ...int) int {
	var min int
	if len(nums) == 0 {
		return 0
	}
	for i, n := range nums {
		if n <= min || i == 0{
			min = n
		}
	}

	return min
}
