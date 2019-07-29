package main

func StrJoin(sep string, strs ...string) string {
	var str string
	lens := len(strs)
	if lens == 0 {
		return str
	}
	for i, s := range strs {
		str += s
		if i != lens-1 {
			str += sep
		}
	}

	return str
}
