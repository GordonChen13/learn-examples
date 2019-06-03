package main

import (
	"fmt"
	"net"
)

func IsUp(v net.Flags) bool {
	return v&net.FlagUp == net.FlagUp
}

func TurnDown(v *net.Flags)  {
	*v &^= net.FlagUp
}

func SetBroadcast(v *net.Flags) {
	*v |= net.FlagBroadcast
}

func IsCast(v net.Flags) bool {
	return v&(net.FlagBroadcast|net.FlagMulticast) != 0
}

func main() {
	var v = net.FlagMulticast | net.FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}
