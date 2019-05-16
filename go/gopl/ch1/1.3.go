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

	if err != nil {
		log.Fatal("args 2 must be number")
	}
	for i := 0; i < count; i++ {
		sa = append(sa, repeat)
	}

    s1 := ""
	before := time.Now().Nanosecond()
	for _, str := range sa {
		s1 += str + " "
	}
	consume1 := time.Now().Nanosecond() - before
	fmt.Println(consume1)

	s2 := ""
	before2 := time.Now().Nanosecond()
	s2 += strings.Join(sa, " ")
	consume2 := time.Now().Nanosecond() - before2
	fmt.Println(consume2)
}
