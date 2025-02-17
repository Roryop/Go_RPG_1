package game

import (
	"fmt"
)

// Gets world_barrier
// Player Chooses World
// Returns world
func Chooseworld(world_barrier int) int {

	var world int
	var isInLoop bool
	for isInLoop {
		fmt.Println("In welche Welt willst du?") // This example is for User deciding world
		fmt.Scanln(&world)

		switch {
		case world == 1 && world_barrier >= 1: // Namen der Welten bitte Eintragen (hab erstmal 3 geschrieben)
			isInLoop = false
		case world == 2 && world_barrier >= 2: // Den welten fehlen vielleicht Attribute aber KA wie mann die schreibt und ob die überhaupt hier sein sollen
			isInLoop = false
		case world == 3 && world_barrier >= 3:
			isInLoop = false
		}
	}
	fmt.Println("Du hast", world, "ausgewählt")
	return world
}
