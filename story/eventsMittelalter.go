package story

import (
	"start/game"
	"start/gear"
	"start/player"
	"start/text"
)

////////////////Funktionen für Wald///////////////////////////

func Baerangriff(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int, typ int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem Räuber überfallen!!!")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, typ)

	return player1, inventory, hp, att, def, rec, world_barrier
}

func Bonfire(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int, typ int) int {
	var maxHp, _, _, _ = player1.CreateStats(inventory)

	text.Print("Du findest einen Lagerfeuer und ruhst dich aus")

	hp = hp + (maxHp / 5)

	return hp
}

////////////////Funktionen für Burg///////////////////////////

////////////////Funktionen für Dorf///////////////////////////
