package story

import (
	"fmt"
	"start/player"
	"start/text"
)

//////////////////////Funktionen für Bar///////////////////////////

// Spieler wählt aus ob er steine an den Rotzbuben zurückwirft
// Gets PlayerStats
// returns nothing
func Whiskey(player1 *player.Player) {
	var choice int

	text.Print("Dir will ein dir unbekannter Einwohner Whiskey ausgeben.")

	// Auswahl Bewohner helfen oder nicht
	text.Print("Nimmst du an?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Der Unbekannte scheint dir sehr charismatisch.")
		text.Print("Whiskey Gläser bringen euch zusammen und ihr freundet euch an.")
		player1.UpdateKarma(7)
	} else {
		text.Print("Aus Höflichkeit oder aus dem Grund, dass der Unbekannte unerträglich stank. Lehnst du ab.")
	}
}

////////////////Funktionen für Shertrif Büro///////////////////////////

////////////////Funktionen für Goldmine///////////////////////////
