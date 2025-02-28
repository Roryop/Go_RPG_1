package menu

import (
	"fmt"
)

// Gets Nothing
// Asks Player what he wants to do
// Returns Choice as Integer
func Menu() int {
	var menuChoice int

	fmt.Println("---------------------------------------------")
	fmt.Println("| 1 - Eine Welt erkunden                    |")
	fmt.Println("| 2 - Aktuellen Schwierigkeitsgrad sehen    |")
	fmt.Println("| 3 - Stats sehen                           |")
	fmt.Println("| 4 - Inventar sehen                        |")
	fmt.Println("| 5 - Map sehen                             |")
	fmt.Println("| 6 - Spiel beenden                         |")
	fmt.Println("---------------------------------------------")

	fmt.Scanln(&menuChoice)

	return menuChoice
}
