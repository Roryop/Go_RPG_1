package story

import (
	"fmt"
	"start/game"
	"start/gear"
	"start/player"
	"start/text"
)

////////////////Funktionen für Wald///////////////////////////

// Spieler wird vom Bär angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Baerangriff(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem Bär angegriffen")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 3)

	return player1, inventory, hp, att, def, rec, world_barrier
}

// Speiler heilt am Lagerfeuer
// Gets PlayerHp, PlayerMaxHp
// Returns PlayerHp mit addiertem ein fünftel vom PlayerMaxHp
func Bonfire(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) int {
	var maxHp, _, _, _ = player1.CreateStats(inventory)

	text.Print("Du findest einen Lagerfeuer und ruhst dich aus")

	hp = hp + (maxHp / 5)

	return hp
}

// Spieler wählt aus ob er einen Pilz isst
// Gets PlayerHP, PlayerMaxHp
// Returns PlayerHp nach essen oder nicht essen vom Pilz
func GesunderPilz(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) int {
	var maxHp, _, _, _ = player1.CreateStats(inventory)
	var choice int

	text.Print("Du findest einen Pilz")

	// Auswahl Pilz essen oder nicht
	text.Print("Möchtest du ihn essen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		hp = hp + (maxHp / 7)

		text.Print("Du hast den Pilz gegessen und etwas regeneriert")
	} else {
		text.Print("Du hast den Pilz nicht gegessen")
	}
	return hp
}

////////////////Funktionen für Burg///////////////////////////

// Spieler wählt aus ob er an Hexenverbrennung teil nimmt
// Gets PlayerStats
// returns nothing
func Hexenjagd(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) {
	var choice int

	text.Print("Es findet eine Hexenverbrennung statt.")

	// Auswahl Pilz essen oder nicht
	text.Print("Möchtest du beim Holz nachlegen helfen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Du spührst wie deine Sündhafte tat deine Wirbelsaüle friert.")
	} else {
		text.Print("Du entscheidest dich nicht in den natürlichen Prozess einzumischen.")
	}
	///////////////////////////////////////////////
	///////////////////////////////////////////////
	/////////            KARMA           //////////
	///////////////////////////////////////////////
	///////////////////////////////////////////////
}

// Spieler kämpft gegen Burgritter
// Gets PlayerStats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Burgritter(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Ein Burgritter fordert dich zu einer Duell heraus.")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 5)

	return player1, inventory, hp, att, def, rec, world_barrier
}

// Spieler wird in den Kerker geworfen
func Kerker(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int
	var maxHp, maxAtt, _, _ = player1.CreateStats(inventory)

	text.Print("Du würdest aufgrund eines Verdachts in den Kerker geworfen.")

	// Auswahl Pilz essen oder nicht
	text.Print("Möchtest du versuchen auszubrechen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Du brichst dir 3 Zehe und 2 Finger.")
		hp = hp - (maxHp / 7)

	} else {
		att = att + (maxAtt / 5)
		text.Print("Durch intensives Training und Meditation baust du Masse auf.")
	}
	return player1, inventory, hp, att, def, rec, world_barrier
}

////////////////Funktionen für Dorf///////////////////////////

// Spieler kämpft gegen Dorfbewohner
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kanpf
func betrunkenerDorfbewohner(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem betrunkenen Dorfbewohner angegriffen!")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 6)

	return player1, inventory, hp, att, def, rec, world_barrier
}
