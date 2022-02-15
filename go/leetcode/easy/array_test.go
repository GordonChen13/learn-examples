package easy

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1,1,2,2,3,4,5,5}
	expect := []int{1,2,3,4,5}
	actual, l := RemoveDuplicates(nums)

	if l != len(expect) {
		t.Fatal("未返回预期长度")
	}
	for i := 0; i < l; i ++ {
		if actual[i] != expect[i] {
			t.Errorf("数组值不相等, 下标: %d, 预期值: %d, 实际值: %d", i, expect[i], actual[i])
		}
	}
}

func TestMaxProfit(t *testing.T) {
	cs :=  []struct{
		Input []int
		Expect int
	}{
		{
			Input: []int{7,1,5,3,6,4},
			Expect: 7,
		},
		{
			Input: []int{1,2,3,4,5},
			Expect: 4,
		},
		{
			Input: []int{7,6,4,3,1},
			Expect: 0,
		},
	}
	for _,c := range cs {
		actual := MaxProfit(c.Input)
		if actual != c.Expect {
			t.Errorf("实际值与预期不符，输入: %v，预期输出: %d, 实际输出: %d", c.Input, c.Expect, actual)
		}
	}
}

func TestRotate(t *testing.T) {
	cs := []struct{
		InputArr []int
		K int
		Expect []int
	}{
		{
			InputArr: []int{1,2,3,4,5,6,7},
			K: 3,
			Expect: []int{5,6,7,1,2,3,4},
		},
		{
			InputArr: []int{1,2,3,4,5,6,7},
			K: 1,
			Expect: []int{7,1,2,3,4,5,6},
		},
		{
			InputArr: []int{-1,-100,3,99},
			K: 2,
			Expect: []int{3,99,-1,-100},
		},
	}

	for _, c := range cs {
		actual := Rotate(c.InputArr, c.K)
		if !IsArrayEqual(c.Expect, actual) {
			t.Errorf("输入: %v, %d, 预期输出: %v, 实际输出: %v", c.InputArr, c.K, c.Expect, actual)
		}
	}
}

func IsArrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestContainsDuplicate(t *testing.T) {
	cs := []struct{
		Input []int
		Expect bool
	}{
		{
			Input: []int{1,2,3,1},
			Expect: true,
		},
		{
			Input: []int{1,2,3,4},
			Expect: false,
		},
		{
			Input: []int{1,1,1,3,3,4,3,2,4,2},
			Expect: true,
		},
	}

	for _, c := range cs {
		actual := ContainsDuplicate(c.Input)
		if actual != c.Expect {
			t.Errorf("输出不符合预期，输入: %v, 预期输出: %v, 实际输出: %v", c.Input, c.Expect, actual)
		}
	}
}

func TestSingleNumber(t *testing.T) {
	cs := []struct{
		Input []int
		Expect int
	} {
		{
			Input:[]int{2,2,1},
			Expect: 1,
		},
		{
			Input:[]int{4,1,2,1,2},
			Expect: 4,
		},
	}

	for _, c := range cs {
		actual := SingleNumber(c.Input)
		if actual != c.Expect {
			t.Errorf("输出不符合预期，输入: %v, 预期输出: %v, 实际输出: %v", c.Input, c.Expect, actual)
		}
	}
}