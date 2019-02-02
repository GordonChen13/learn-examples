package npcs

import (
	"fmt"
	"math"
)

func (loc Location) String() string{
	return fmt.Sprintf("(%f,%f,%f)", loc.X, loc.Y, loc.Z)
}

func (loc Location) EnclideanDistance(target Location) float64 {
	return math.Sqrt(
		(loc.X - target.X)*(loc.X - target.X) + (loc.Y - target.Y)*(loc.Y - target.Y) + (loc.Z - target.Z)*(loc.Z - target.Z))
}

func (npc NonPlayerCharacter) DistanceTo(target NonPlayerCharacter) float64  {
	return npc.Loc.EnclideanDistance(target.Loc)	
}

func (npc NonPlayerCharacter) String() string {
	return fmt.Sprintf("%s %s", npc.Name, npc.Loc)
}
