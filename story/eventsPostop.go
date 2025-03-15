package story

import (
	"fmt"
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

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 7)

	return player1, inventory, hp, att, def, rec, world_barrier
}

// Spieler wird von einer dreikoepfigen Schlange angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Dreikoepfigeschlange(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von eine dreikoepfigen Schlange angegriffen!")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 8)

	return player1, inventory, hp, att, def, rec, world_barrier
}

// Spieler wird von einer laufenden Makrowelle angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Laufendemakrowelle(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einer laufenden Makrowelle angegriffen!")

	player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 9)

	return player1, inventory, hp, att, def, rec, world_barrier
}

////////////////Funktionen für Siedlung///////////////////////////

// Spieler wird von einer laufenden Makrowelle angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Plünderer(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int

	text.Print("Die Siedlung wird von Plünderern überrannt.")

	// Auswahl Pilz essen oder nicht
	text.Print("Willst du bei der Verteidigung helfen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {

		text.Print("Du wirst von einem Plünderer angegriffen!")
		player1, inventory, hp, att, def, rec, world_barrier = game.Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 10)

	} else {
		text.Print("Du bist wie ein Feige weggerannt")
	}

	return player1, inventory, hp, att, def, rec, world_barrier
}

// Spieler wählt aus ob er einen Pilz isst
// Gets PlayerHP, PlayerMaxHp
// Returns PlayerHp nach essen oder nicht essen vom Pilz
func Stammarzt(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) int {
	var maxHp, _, _, _ = player1.CreateStats(inventory)
	var choice int

	text.Print("Der Stammarzt möchte ein neues Medikament ausprobieren.")

	// Auswahl Pilz essen oder nicht
	text.Print("Willst du es testen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		hp = hp + (maxHp / 2)

		text.Print("Du spührst wie Kreatinphosphat deine Muskeln mit Wasser füllt!!!")
	} else {
		text.Print("Du wolltest lieber ein angebissener Lauch bleiben")
	}
	return hp
}

////////////////////Funktionen für Geisterstdt///////////////////////////
