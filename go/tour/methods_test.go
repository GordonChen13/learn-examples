package tour

import (
	"testing"
	"fmt"
)

func TestIPAddr_String(t *testing.T) {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func TestLoopSqrtError(t *testing.T) {
	x := -10.0
	y := 100
	LoopSqrtError(x, y)
}