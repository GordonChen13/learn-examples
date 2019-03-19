package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const finalWord = "go!"
const countdownStart  = 3

func main()  {
	Countdown(os.Stdout)
}

func Countdown(writer io.Writer)  {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(writer, strconv.Itoa(i))
	}
	time.Sleep(1 * time.Second)
	fmt.Fprintf(writer, finalWord)
}
