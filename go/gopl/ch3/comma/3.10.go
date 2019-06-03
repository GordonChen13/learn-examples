package comma

import "bytes"

func CommaB(s string) string {
	var buf bytes.Buffer
	n := len(s)
	l := n%3
	if l == 0 {
		l  = 3
	}
	buf.WriteString(s[:l])
	s = s[l:]
	for len(s) >= 3 {
       	buf.WriteString("," + s[:3])
        s = s[3:]
	}
	buf.WriteString(s)

	return buf.String()
}
