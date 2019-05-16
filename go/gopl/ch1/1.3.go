package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main()  {
	repeat := os.Args[1]
	count, err := strconv.Atoi(os.Args[2])
	sa := make([]string, count)
	var s string

	if err != nil {
		log.Fatal("args 2 must be number")
	}
	for i := 0; i < count; i++ {
		sa = append(sa, repeat)
	}

	before := time.Now().Nanosecond()
	for _, str := range sa {
		s = str + " "
	}
	consume1 := time.Now().Nanosecond() - before
	fmt.Println(consume1)

	s = ""
	before2 := time.Now().Nanosecond()
	s = strings.Join(sa, " ")
	consume2 := time.Now().Nanosecond() - before2
	fmt.Println(consume2)
}
