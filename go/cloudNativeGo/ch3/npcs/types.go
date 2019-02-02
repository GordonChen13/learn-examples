package npcs

type Power struct {
	Attack int
	Defence int
}

type Location struct {
	X float64
	Y float64
	Z float64
}

type NonPlayerCharacter struct {
	Name string
	Speed int
	HP int
	Power Power
	Loc Location
}
