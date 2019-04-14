package poker

import (
	"fmt"
	"io"
	"os"
	"time"
)

// BlindAlerter schedules alerts for blind amounts
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

// BlindAlerterFunc allows you to implement BlindAlerter with a function
type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

// ScheduleAlertAt is BlindAlerterFunc implementation of BlindAlerter
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	a(duration, amount, to)
}

// StdOutAlerter will schedule alerts and print them to os.Stdout
func StdOutAlerter(duration time.Duration, amount int, to io.Writer) {
	to = os.Stdout
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}
