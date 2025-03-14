package story

import (
	"fmt"
	"start/game"
	"start/gear"
	"start/player"
	"start/text"
	"strings"
)

// Gets player, inventory, current stats, world, world_barrier
// Executes Event Robbery
// Returns player, inventory, current  stats, world_barrier
///////////////////Cyberpunk Welt////////////////////////
//////////////////Funktionen für Slums///////////////////

func Robbery(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem Räuber überfallen!!!")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 4)

	return player1, inventory, hp, att, def, rec, world_barrier
}

func Bettler(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var giveaway string
	text.Print("Du triffst auf einen Bettler. Willst du ihm etwas geben? /ja /nein?")
	fmt.Scanln(&giveaway) //////////Entscheidung des Spielers ob er eas abgibt oder nicht
	giveaway = strings.ToLower(giveaway)

	switch {
	case giveaway == "ja": //////Er gibt etwas ab -> Der Penner dankt ihm und der Spieler erlangt Karma
		text.Print("Du gibst dem Penner eins deiner Items und er dankt dir.")
		/////////////Itemdrop hier
		/////////////Karma++
	case giveaway == "nein":
		text.Print("Der Obdachlose wird sauer und greift dich an!")
		player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 18)
	}

	return player1, inventory, hp, att, def, rec, world_barrier
}

//////////////////Funktionen für Businessviertel///////////////////

//////////////////Funktionen für Außenstadt///////////////////
