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
		player1.UpdateKarma(6)
	} else {
		text.Print("Du bist geflüchtet.")
		text.Print("...")
		text.Print("Der Anführer der Bande hatte nicht mal ein Pferd.")
	}

	return player1, inventory, hp, att, def, rec, player_level
}

////////////////Funktionen für Sheriff Büro///////////////////////////

// Spieler wählt aus ob er einen Insassen Freilässt
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Jailbreak(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int
	var item = gear.ItemDrop(player_level)
	var gearTyp = item.GetGearTyp()

	text.Print("Der Insasse des Gefängnisses bittet dich heimlich, ihn freizulassen.")

	// Auswahl Freilassen oder nicht
	text.Print("Willst du auf ihn Hören?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Der Insasse bedankt sich bei dir.")
		if gearTyp == "Empty" {
			text.Print("Jedoch aber nur wörtlich.")
		} else {
			inventory = gear.AddToInventory(inventory, item)
		}
		text.Print("Der Sheriff riecht aber dein Missbrauch von Freiheit von 3/4 Meile aus.")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 16)
	} else {
		text.Print("Du lässt ihn in der Zelle gammeln.")
		player1.UpdateKarma(4)
	}

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler wählt aus ob er auf einen Insassen aufpassen will
// Gets inventory (fügt einen Item hinzu falls der Spieler dem Sheriff hilft)
// returns inventory
func Sheriffquest(player1 *player.Player, inventory [10]*gear.InventorySlot, player_level int) [10]*gear.InventorySlot {
	var choice int
	var item = gear.ItemDrop(player_level)

	text.Print("Der Sherrif fragt dich, ob du kurz auf einen gefangenen aufpassen kannst.")

	// Auswahl plündern oder nicht
	text.Print("Willst du dir die Zeit nehmen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Der Sheriff schenkt dir eine Waffe aus seinem Arsenal.")
		inventory = gear.AddToInventory(inventory, item)
	} else {
		text.Print("Würd mir stinken...")
	}
	return inventory
}

////////////////Funktionen für Goldmine///////////////////////////
