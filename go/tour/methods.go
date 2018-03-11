package tour

import "fmt"

type IPAddr [4]byte
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

type SqrtError float64

func (e SqrtError) Error() string  {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func LoopSqrtError(s SqrtError, n int) (float64, e error)  {
	if s < 0 {
		e := "error"
	}
	var z SqrtError
	z = 1.0
	for i := 0; i < n; i ++ {
		z -= (s * s - z) / 2*z
	}
}
