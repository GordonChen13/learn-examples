package easy

import "testing"

func TestReverseString(t *testing.T) {
	cs := []struct {
		Input  []byte
		Expect []byte
	}{
		{
			Input:  []byte{'h', 'e', 'l', 'l', 'o'},
			Expect: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			Input:  []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			Expect: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, c := range cs {
		ReverseString(c.Input)
		for i, v := range c.Input {
			if c.Expect[i] != v {
				t.Fatalf("与预期结果不符，实际输出: %s, 预期输出: %s", c.Input, c.Expect)
			}
		}
	}
}
