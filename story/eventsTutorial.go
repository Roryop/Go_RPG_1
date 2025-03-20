package story

import (
	"fmt"
	"start/gear"
	"start/player"
	"start/text"
)

// Spieler wird vom Ork angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Ork(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem Ork angegriffen")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 1)

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler lernt karma kennen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Demokarma(player1 *player.Player) {
	var choice int

	text.Print("Ein Alter Mann nähert sich...")
	text.Print("Er bietet dir magische Bohnen an.")

	// Auswahl Pilz essen oder nicht
	text.Print("Nimmst du die?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Die Schmecken nach Socken aber du warst nett zu dem alten Kerl.")
		player1.UpdateKarma(3)
	} else {
		text.Print("Der Mann ist traurig...")
		text.Print("...")
		text.Print("Fühlst du dich jetzt besser?")
		player1.UpdateKarma(-3)
	}
}
