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

func TestIntersect(t *testing.T) {
	cs := []struct{
		Input1 []int
		Input2 []int
		Expect []int
	} {
		{
			Input1:[]int{1,2,2,1},
			Input2:[]int{2,2},
			Expect: []int{2,2},
		},
		{
			Input1:[]int{4,9,5},
			Input2:[]int{9,4,9,8,4},
			Expect: []int{4,9},
		},
		{
			Input1:[]int{1,3,8,9,3},
			Input2:[]int{1, 0},
			Expect: []int{1},
		},
	}

	for _, c := range cs {
		actual := Intersect(c.Input1, c.Input2)
		if !IsArrayEqual(actual, c.Expect) {
			t.Errorf("输出不符合预期，输入: %v, %v, 预期输出: %v, 实际输出: %v", c.Input1, c.Input2, c.Expect, actual)
		}
	}
}

func TestPlusOne(t *testing.T) {
	cs := []struct{
		Input  []int
		Expect []int
	} {
		{
			Input:  []int{1,2,3},
			Expect: []int{1,2,4},
		},
		{
			Input:  []int{4,3,2,1},
			Expect: []int{4,3,2,2},
		},
		{
			Input:  []int{0},
			Expect: []int{1},
		},
		{
			Input:  []int{9,9,9},
			Expect: []int{1,0,0,0},
		},
	}

	for _, c := range cs {
		actual := PlusOne(c.Input)
		if !IsArrayEqual(actual, c.Expect) {
			t.Errorf("输出不符合预期，输入: %v, 预期输出: %v, 实际输出: %v", c.Input, c.Expect, actual)
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	cs := []struct{
		Input  []int
		Expect []int
	} {
		{
			Input:  []int{0,1,0,3,12},
			Expect: []int{1,3,12,0,0},
		},
		{
			Input:  []int{0},
			Expect: []int{0},
		},
		{
			Input:  []int{0,1,0,0},
			Expect: []int{1,0,0,0},
		},
		{
			Input:  []int{0,0,1},
			Expect: []int{1,0,0},
		},
		{
			Input:  []int{1,0,1},
			Expect: []int{1,1,0},
		},
		{
			Input:  []int{0,0,0,0},
			Expect: []int{0,0,0,0},
		},
	}

	for _, c := range cs {
		MoveZeroes(c.Input)
		if !IsArrayEqual(c.Input, c.Expect) {
			t.Errorf("输出不符合预期，预期输出: %v, 实际输出: %v", c.Expect, c.Input)
		}
	}
}

func TestTwoSum(t *testing.T) {
	cs := []struct{
		Nums   []int
		Target int
		Expect []int
	} {
		{
			Nums:   []int{2,7,11,15},
			Target: 9,
			Expect: []int{0,1},
		},
		{
			Nums:   []int{3,2,4},
			Target: 6,
			Expect: []int{1,2},
		},
		{
			Nums:   []int{3,3},
			Target: 6,
			Expect: []int{0,1},
		},
	}

	for _, c := range cs {
		actual := TwoSum(c.Nums, c.Target)
		if !IsArrayEqual(actual, c.Expect) {
			t.Errorf("输出不符合预期，输入: %v, %v, 预期输出: %v, 实际输出: %v", c.Nums, c.Target, c.Expect, actual)
		}
	}
}

func TestIsValidSudoku(t *testing.T) {
	cs := []struct{
		Board [][]byte
		Expect bool
	} {
		{
			Board: [][]byte{
				{'5','3','.','.','7','.','.','.','.'},
				{'6','.','.','1','9','5','.','.','.'},
				{'.','9','8','.','.','.','.','6','.'},
				{'8','.','.','.','6','.','.','.','3'},
				{'4','.','.','8','.','3','.','.','1'},
				{'7','.','.','.','2','.','.','.','6'},
				{'.','6','.','.','.','.','2','8','.'},
				{'.','.','.','4','1','9','.','.','5'},
				{'.','.','.','.','8','.','.','7','9'},
			},
			Expect: true,
		},
		{
			Board: [][]byte{
				{'8','3','.','.','7','.','.','.','.'},
				{'6','.','.','1','9','5','.','.','.'},
				{'.','9','8','.','.','.','.','6','.'},
				{'8','.','.','.','6','.','.','.','3'},
				{'4','.','.','8','.','3','.','.','1'},
				{'7','.','.','.','2','.','.','.','6'},
				{'.','6','.','.','.','.','2','8','.'},
				{'.','.','.','4','1','9','.','.','5'},
				{'.','.','.','.','8','.','.','7','9'},
			},
			Expect: false,
		},
		{
			Board: [][]byte{
				{'8','3','.','.','7','.','.','.','.'},
				{'6','.','.','1','9','5','.','.','.'},
				{'.','9','8','.','.','.','.','6','.'},
				{'1','.','.','.','6','.','.','.','3'},
				{'4','.','.','8','.','3','.','.','1'},
				{'7','.','.','.','2','.','.','.','6'},
				{'.','6','.','.','.','.','2','8','.'},
				{'.','.','.','4','1','9','.','.','5'},
				{'.','.','.','.','8','.','.','7','9'},
			},
			Expect: false,
		},
	}

	for _, c := range cs {
		actual := IsValidSudoku(c.Board)
		if actual != c.Expect {
			t.Errorf("输出不符合预期，输入: %v, 预期输出: %v, 实际输出: %v", c.Board, c.Expect, actual)
		}
	}
}
