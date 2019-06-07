package sort

func Insert(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		j := i-1
		for ; j >= 0 && nums[j] > x; j-- {
			nums[j+1] = nums[j]
			nums[j] = x
		}
	}

	return nums
}
