package story

import (
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
