package easy

import (
	"sort"
)

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

// Intersect 两个数组的交集
// @see https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/
func Intersect(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var result []int
	var i, j int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

// PlusOne 加一
// @see https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2cv1c/
func PlusOne(nums []int) []int {
	for i := len(nums)-1; i >= 0; i-- {
		if nums[i] != 9 {
			nums[i] = nums[i] + 1
			return nums
		} else {
			nums[i] = 0
		}
	}

	var result []int
	result = append(result, 1)
	for _, num := range nums {
		result = append(result, num)
	}

	return result
}

// MoveZeroes 移动零
// @see https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2ba4i/
func MoveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// TwoSum 两数之和
// @see https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2jrse/
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var j int
	var ok bool
	for i, v := range nums {
		j, ok = m[target - v]
		if ok {
			return []int{j,i}
		}
		m[v] = i
	}
	return []int{}
}

// IsValidSudoku 有效的数独
// @see https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2f9gg/
func IsValidSudoku(board [][]byte) bool {
	var rowCount, colCount [9][9]int
	var subboxes [3][3][9]int

	for i, nums := range board {
		for j, numB := range nums {
			if numB == '.' {
				continue
			}
			num := numB - '1'
			rowCount[i][num]++
			colCount[j][num]++

			subboxes[i/3][j/3][num]++

			if rowCount[i][num] > 1 || colCount[j][num] > 1 || subboxes[i/3][j/3][num] > 1{
				return false
			}
		}
	}
	return true
}