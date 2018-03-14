package tour

import (
	"testing"
	"fmt"
	"strings"
	"io"
	"os"
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

func TestRot13Reader_Read(t *testing.T) {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}