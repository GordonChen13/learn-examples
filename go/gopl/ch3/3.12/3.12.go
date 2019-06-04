package main

import "strings"

func HasSameChar(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		find := strings.LastIndex(s2, string(s1[i]))
		if find < 0 {
			return false
		}
		s2 = s2[:find] + s2[find+1:]
	}
	if len(s2) > 0 {
		return false
	}
	return true
}
