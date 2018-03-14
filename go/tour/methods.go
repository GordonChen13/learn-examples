package tour

import (
	"fmt"
	"io"
)

type IPAddr [4]byte
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

type SqrtError float64

func (e SqrtError) Error() string  {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	}
	return "hah"
}

func LoopSqrtError(s float64, n int) (float64, error)  {
	if s < 0 {
		return 0, SqrtError(s)
	}
	z := 1.0
	for i := 0; i < n; i ++ {
		z -= (s * s - z) / 2*z
	}
	return  z, nil
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.


type rot13Reader struct {
	r io.Reader
}
