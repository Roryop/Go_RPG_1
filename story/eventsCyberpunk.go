package story

import (
	"start/game"
	"start/gear"
	"start/player"
	"start/text"
)

// Gets player, inventory, current stats, world, world_barrier
// Executes Event Robbery
// Returns player, inventory, current  stats, world_barrier
func Robbery(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int, typ int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem Räuber überfallen!!!")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, typ)

	return player1, inventory, hp, att, def, rec, world_barrier
}
