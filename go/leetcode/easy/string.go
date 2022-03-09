package easy

// ReverseString 反转字符串
// @see https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhbqj/
func ReverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		temp := s[i]
		s[i] = s[l-i-1]
		s[l-i-1] = temp
	}
}
