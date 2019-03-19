package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{0}

	Countdown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}