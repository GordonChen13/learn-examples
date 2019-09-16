package main

import (
	"fmt"
	"syscall"
)

type Errono unintptr

ar errors = [...] string{
1:   "operation not permitted",
2:   "no such file or directory",
3:   "no such process",
}

var (e Errono) Error() String {
if 0 <= int(e) && int(e) < len(errors) {
return errors[e]
}
return fmt.Sprintf("errno %d", e)
}

func main() {
	var err error = syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}
