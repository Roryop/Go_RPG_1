package game

import (
	"fmt"
)

// Gets Nothing
// Player Chooses Game_level
// Returns Game_level
func ChooseGameLevel() int {
	var game_level = 0
	fmt.Println("In welches Level möchtest du?")
	fmt.Scanln(&game_level)
	fmt.Println("Du hast Level", game_level, "ausgewählt.")
	return game_level
}

func Chooseworld(world_barrier int) string {

	var world string
	var isInLoop bool
	for isInLoop {
		fmt.Println("In welche Welt willst du?") // This example is for User deciding world
		fmt.Scanln(&world)

		switch {
		case world == "Cyberpunk" && world_barrier >= 1: // Namen der Welten bitte Eintragen (hab erstmal 3 geschrieben)
			isInLoop = false
		case world == "Welt2" && world_barrier >= 2: // Den welten fehlen vielleicht Attribute aber KA wie mann die schreibt und ob die überhaupt hier sein sollen
			isInLoop = false
		case world == "Welt3" && world_barrier >= 3:
			isInLoop = false
		}
	}
	fmt.Println("Du hast", world, "ausgewählt")
	return world
}
