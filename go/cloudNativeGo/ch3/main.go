package main

import (
	"fmt"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch3/npcs"
)

func main()  {
	mob := npcs.NonPlayerCharacter{Name:"Alfred"}
	fmt.Println(mob)

	hortense := npcs.NonPlayerCharacter{Name:"Hortense", Loc:npcs.Location{X: 10.0, Y:10.0, Z:10.0}}
	fmt.Println(hortense)

	fmt.Printf("Alfred is %f units from Hortense.\n", mob.DistanceTo(hortense))
}
