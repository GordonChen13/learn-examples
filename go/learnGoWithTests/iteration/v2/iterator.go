package v1

import (
	"fmt"
)

const repeatCount  = 5

func Repeat(a string) string {
	var repeated string
	for i := 0; i < repeatCount; i++{
		repeated = repeated + a
	}
	return repeated
}

func main()  {
	fmt.Println(Repeat("a"))
}
