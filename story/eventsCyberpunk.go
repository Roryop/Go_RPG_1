package story

import (
	"fmt"
	"start/gear"
	"start/player"
	"start/text"
)

///////////////////Cyberpunk Welt////////////////////////

//////////////////Funktionen für Slums///////////////////

// Gets player, inventory, current stats, world, world_barrier
// Executes Event Robbery
// Returns player, inventory, current  stats, world_barrier
func Robbery(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem Räuber überfallen!!!")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 4)

	return player1, inventory, hp, att, def, rec, player_level
}

func Bettler(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var giveaway int
	text.Print("Du triffst auf einen Bettler. Willst du ihm etwas geben? /ja /nein?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&giveaway) //////////Entscheidung des Spielers ob er eas abgibt oder nicht

	switch {
	case giveaway == 1: //////Er gibt etwas ab -> Der Penner dankt ihm und der Spieler erlangt Karma
		text.Print("Du gibst dem Penner eins deiner Items und er dankt dir.")
		gear.SubtractFromInventory(inventory)

		player1.UpdateKarma(7)
	case giveaway == 2:
		text.Print("Der Obdachlose wird sauer und greift dich an!")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 18)
	}

	return player1, inventory, hp, att, def, rec, player_level
}

func Muelltonne(inventory [10]*gear.InventorySlot, player_level int) ([10]*gear.InventorySlot, int) {
	var choice int
	text.Print("Du läufst an einer Mülltonne vorbei, willst du sie durchsuchen? /ja /nein?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	switch {
	case choice == 1: ///////////////In der Tonne liegt ein Item
		text.Print("Du wühlst in einer Mülltonne herum. Denk mal drüber nach!")
		var item = gear.ItemDrop(player_level)
		gear.AddToInventory(inventory, item)

	case choice == 2: ////////////////Nix passiert
		text.Print("OK dann eben nicht")

	}
	return inventory, player_level
}

//////////////////Funktionen für Businessviertel///////////////////

//////////////////Funktionen für Außenstadt///////////////////

func Grenzkontrolle(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	text.Print("Du gerätst in eine Grenzkontrolle. Der Polizist ist schlecht gelaunt und greift Dich an")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 20)

	return player1, inventory, hp, att, def, rec, player_level
}
