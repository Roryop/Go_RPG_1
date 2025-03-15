package story

import (
	"start/game"
	"start/gear"
	"start/player"
	"start/text"
)

/////////////////////Funktionen für Ground Zero///////////////////////////

// Spieler wird von einer menschengroßen Mutantenratte angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Mutantenratte(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einer Menschengroßen Mutantenratte angegriffen!")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 6)

	return player1, inventory, hp, att, def, rec, world_barrier
}

// Spieler wird von einer dreikoepfigen Schlange angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Dreikoepfigeschlange(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von eine dreikoepfigen Schlange angegriffen!")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 6)

	return player1, inventory, hp, att, def, rec, world_barrier
}

// Spieler wird von einer laufenden Makrowelle angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Laufendemakrowelle(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einer laufenden Makrowelle angegriffen!")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 6)

	return player1, inventory, hp, att, def, rec, world_barrier
}

////////////////////Funktionen für Geisterstdt///////////////////////////
