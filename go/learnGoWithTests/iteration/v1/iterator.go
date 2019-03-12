package v1

import (
	"fmt"
)

func Repeat(a string) string {
	var repeated string
	for i := 0; i < 5; i++{
		repeated = repeated + a
	}
	return repeated
}

func main()  {
	fmt.Println(Repeat("a"))
}
