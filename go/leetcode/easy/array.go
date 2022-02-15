package easy

// RemoveDuplicates 删除排序数组中的重复项
// @see https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
func RemoveDuplicates(nums []int) ([]int, int) {
	var i,j int
	for j = 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return nums, i+1
}

// MaxProfit 买卖股票的最佳时机
// @see https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
func MaxProfit(prices []int) int {

	return 0
}

// Rotate 旋转数组
// @see https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
func Rotate(nums []int, k int) []int  {
	l := len(nums)
	if l == 1 {
		return nums
	}
	if k > l {
		k = k%l
	}
	cp := make([]int, l)
	copy(cp, nums)
	for i, v := range cp {
		if i + k < l {
			nums[i+k]	= v
		} else {
			nums[i+k-l] = v
		}
	}

	return nums
}

// ContainsDuplicate 存在重复元素
// @see https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x248f5/
func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, n := range nums {
		if m[n] {
			return true
		}
		m[n] = true
	}
	return false
}

// SingleNumber 只出现一次的数字
// @see https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
func SingleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result = result  ^ n
	}
	return result
}