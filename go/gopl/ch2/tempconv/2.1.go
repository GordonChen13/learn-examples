package tempconv

import "fmt"

type Kelvin float64

const AbsoluteZeroK Kelvin = -273.15

func (k Kelvin) String() string {
	return fmt.Sprintf("%g K", k)
}

func (k Kelvin) KToC() Celsius {
	return Celsius(k)
}
