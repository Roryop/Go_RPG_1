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
