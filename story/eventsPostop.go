package story

import (
	"fmt"
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

	player1, inventory, hp, att, def, rec, world_barrier = Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 7)

	return player1, inventory, hp, att, def, rec, world_barrier
}

// Spieler wird von einer dreikoepfigen Schlange angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func DreikoepfigeSchlange(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von eine dreikoepfigen Schlange angegriffen!")

	player1, inventory, hp, att, def, rec, world_barrier = Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 8)

	return player1, inventory, hp, att, def, rec, world_barrier
}

// Spieler wird von einer laufenden Makrowelle angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func LaufendeMakrowelle(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einer laufenden Makrowelle angegriffen!")

	player1, inventory, hp, att, def, rec, world_barrier = Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 9)

	return player1, inventory, hp, att, def, rec, world_barrier
}

////////////////Funktionen für Siedlung///////////////////////////

// Spieler wird von einer laufenden Makrowelle angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func PluendererGroup(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int

	text.Print("Die Siedlung wird von Plünderern überrannt.")

	// Auswahl Pilz essen oder nicht
	text.Print("Willst du bei der Verteidigung helfen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {

		text.Print("Du wirst von einem Plünderer angegriffen!")
		player1, inventory, hp, att, def, rec, world_barrier = Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 10)

	} else {
		text.Print("Du bist wie ein Feigling weggerannt")
	}

	return player1, inventory, hp, att, def, rec, world_barrier
}

// Spieler wählt aus ob er einen Pilz isst
// Gets PlayerHP, PlayerMaxHp
// Returns PlayerHp nach essen oder nicht essen vom Pilz
func Stammarzt(player1 *player.Player, inventory [10]*gear.InventorySlot, hp int) int {
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

// Spieler wählt aus ob er steine an den Rotzbuben zurückwirft
// Gets PlayerStats
// returns nothing
func VerletzterBewohner() {
	var choice int

	text.Print("Du triffst einen verletzten Bewohner.")

	// Auswahl Pilz essen oder nicht
	text.Print("Willst du ihm helfen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Du bringst den Bewohner in eine lehre Bude, bindest seine Wunden und gehst.")
		text.Print("Kurz danach siehst du aus einer Ferne wie diese Bude von jemandem in die Luft gesprengt wird.")
		text.Print("...")
		text.Print("Naja... Letzendlich hast du ihm geholfen.")
		// + Karma
	} else {
		text.Print("Du entscheidest dich ihn schmerzhaft sterben zu lassen.")
	}
	///////////////////////////////////////////////
	///////////////////////////////////////////////
	/////////            KARMA -         //////////
	///////////////////////////////////////////////
	///////////////////////////////////////////////
}

////////////////////Funktionen für Geisterstdt///////////////////////////

// Spieler wird vom Pluenderer angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func PluendererSingle(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Auf dich trifft ein Plünderer, er möchte dich berauben!")

	player1, inventory, hp, att, def, rec, world_barrier = Fight(player1, inventory, hp, att, def, rec, world, world_barrier, 10)

	return player1, inventory, hp, att, def, rec, world_barrier
}
