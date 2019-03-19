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

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}
type DefaultSleeper struct {}

func (s *SpySleeper) Sleep()  {
	s.Calls++
}

func (d *DefaultSleeper) Sleep()  {
	time.Sleep(1 * time.Second)
}

func main()  {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(writer io.Writer, s Sleeper)  {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(writer, strconv.Itoa(i))
	}
	s.Sleep()
	fmt.Fprintf(writer, finalWord)
}
