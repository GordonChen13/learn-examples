package v1

import "strings"

func Repeat(a string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++{
		repeated = repeated + a
	}
	return repeated
}

func PkgRepeat(a string, repeatCount int) string {
	repeated := strings.Repeat(a, repeatCount)
	return repeated
}
