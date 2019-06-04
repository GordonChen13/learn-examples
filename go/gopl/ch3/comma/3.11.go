package comma

import "strings"

func CommaC(s string) string {
	var mPart,dotPart,cPart string
	minus := strings.LastIndex(s, "-")
	if minus >= 0 {
		mPart = "-"
		s = s[minus+1:]
	}
	dot := strings.LastIndex(s, ".")
	if dot >= 0 {
		dotPart = s[dot:]
		s = s[:dot]
	}
	cPart = Comma(s)

	return mPart + cPart +  dotPart
}
