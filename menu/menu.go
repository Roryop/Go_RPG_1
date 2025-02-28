package menu

import (
	"fmt"
)

// Gets Nothing
// Asks Player what he wants to do
// Returns Choice as Integer
func Menu() int {
	var menuChoice int

	fmt.Println("-----------------------------------------")
	fmt.Println("| 1 - Explore a World                   |")
	fmt.Println("| 2 - Know your current difficulty      |")
	fmt.Println("| 3 - See your Stats                    |")
	fmt.Println("| 4 - See your Inventory                |")
	fmt.Println("| 5 - See the Map                       |")
	fmt.Println("| 6 - End the Game                      |")
	fmt.Println("-----------------------------------------")

	fmt.Scanln(&menuChoice)

	return menuChoice
}
