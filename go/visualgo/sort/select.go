package sort

func Select(nums []int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
            min := nums[i]
			if nums[i] > nums[j] {
                min = nums[j]
				nums[j] = nums[i]
			}
			nums[i] = min
		}
	}
	return nums
}

