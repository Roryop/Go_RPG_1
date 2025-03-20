package story

import (
	"fmt"
	"start/gear"
	"start/player"
	"start/text"
)

//////////////////////Funktionen für Bar///////////////////////////

// Spieler wählt aus ob er jemanden für Whiskey ausgeben lässt
// Gets PlayerStats
// returns karma
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

// Spieler wählt aus ob er sich in eine Schlaegerei einmischen will
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Schlaegerei(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int

	text.Print("In der Bar bricht eine Schlaegerei aus!")

	// Auswahl Verteidigung helfen oder nicht
	text.Print("Willst du deich einmischen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Es kommt zu einem Kampf eins gegen eins!")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 14)

	} else {
		text.Print("Du hast dich unter dem Bartisch versteckt. Bis alles vorbei ging.")

	}

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler wählt aus ob er sich in eine Schlaegerei einmischen will
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Barüberfall(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int

	text.Print("Die Bar wird überfallen.")

	// Auswahl Verteidigung helfen oder nicht
	text.Print("Willst du die Bar verteidigen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Der Anführer der Bande taucht vor dir auf...")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 15)

	} else {
		text.Print("Du bist geflüchtet.")
		text.Print("...")
		text.Print("Der Anführer der Bande hatte nicht mal ein Pferd.")
	}

	return player1, inventory, hp, att, def, rec, player_level
}

////////////////Funktionen für Shertrif Büro///////////////////////////

////////////////Funktionen für Goldmine///////////////////////////
