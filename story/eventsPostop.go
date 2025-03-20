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
func Mutantenratte(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einer Menschengroßen Mutantenratte angegriffen!")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 7)

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler wird von einer dreikoepfigen Schlange angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func DreikoepfigeSchlange(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von eine dreikoepfigen Schlange angegriffen!")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 8)

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler wird von einer laufenden Makrowelle angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func LaufendeMakrowelle(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einer laufenden Makrowelle angegriffen!")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 9)

	return player1, inventory, hp, att, def, rec, player_level
}

////////////////Funktionen für Siedlung///////////////////////////

// Spieler wird von einer laufenden Makrowelle angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func PluendererGroup(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int

	text.Print("Die Siedlung wird von Plünderern überrannt.")

	// Auswahl Verteidigung helfen oder nicht
	text.Print("Willst du bei der Verteidigung helfen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {

		text.Print("Du wirst von einem Plünderer angegriffen!")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 10)

		text.Print("Du wirst von einem zweiten Plünderer angegriffen!")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 10)

	} else {
		text.Print("Du bist wie ein Feigling weggerannt")

		player1.UpdateKarma(-4)
	}

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler wählt aus ob er den Stammarzt experimentieren lassen will
// Gets Player, inventory, curent stats
// Returns PlayerHp nach essen oder nicht essen vom Medikament
func Stammarzt(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int) int {
	var choice int

	text.Print("Der Stammarzt möchte ein neues Medikament ausprobieren.")

	// Auswahl Medizin testen oder nicht
	text.Print("Willst du es testen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		hp, _, _, _ = player1.UpgradeStat(inventory, hp, att, def, rec, 1)

		text.Print("Du spührst wie Kreatinphosphat deine Muskeln mit Wasser füllt!!!")
	} else {
		text.Print("Du wolltest lieber ein angebissener Lauch bleiben")
	}
	return hp
}

// Spieler wählt aus ob er steine an den Rotzbuben zurückwirft
// Gets PlayerStats
// returns nothing
func VerletzterBewohner(player1 *player.Player) {
	var choice int

	text.Print("Du triffst einen verletzten Bewohner.")

	// Auswahl Bewohner helfen oder nicht
	text.Print("Willst du ihm helfen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Du bringst den Bewohner in eine lehre Bude, bindest seine Wunden und gehst.")
		text.Print("Kurz danach siehst du aus einer Ferne wie diese Bude von jemandem in die Luft gesprengt wird.")
		text.Print("...")
		text.Print("Naja... Letzendlich hast du ihm geholfen.")

		player1.UpdateKarma(6)
	} else {
		text.Print("Du entscheidest dich ihn schmerzhaft sterben zu lassen.")

		player1.UpdateKarma(-5)
	}
}

////////////////////Funktionen für Geisterstdt///////////////////////////

// Spieler wird vom Pluenderer angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func PluendererSingle(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Auf dich trifft ein Plünderer, er möchte dich berauben!")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 10)

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler wählt aus ob er ein Geschäft plündern will
// Gets inventory (fügt einen Item hinzu falls er einen findet)
// returns inventory
func Geschäftplündern(player1 *player.Player, inventory [10]*gear.InventorySlot, player_level int) [10]*gear.InventorySlot {
	var choice int
	var item = gear.ItemDrop(player_level)
	var gearTyp = item.GetGearTyp()

	text.Print("Du findest ein verlassenes Geschäft.")

	// Auswahl plündern oder nicht
	text.Print("Willst du es plündern?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		if gearTyp == "Empty" {
			text.Print("Du hast erfolgreich nichts gefunden.")

		} else {
			text.Print("Du findest legendäre Artefakten!")
			inventory = gear.AddToInventory(inventory, item)
		}
	} else {
		text.Print("Du entgehst der möglichkeit Konsequenzenlos Schatz zu gewinnen")
		text.Print("...")
		text.Print("Würd mir stinken...")
	}
	return inventory
}

// Spieler wählt aus ob er den Hilfeschreien folgt
// Gets PlayerStats, Kampfalgoritmus
// returns karma, playerstats nach dem Kampf
func Hilfeschreie(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int

	text.Print("Du hörst Hilfeschreie")

	// Auswahl folgen oder nicht
	text.Print("Willst du diesen folgen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Du entscheidest dich ausnahmsweise nett zu sein und der schreienden Person zu helfen.")
		player1.UpdateKarma(6)
		text.Print("Die schreihe waren aber eine Imitation von einem Monster!")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 19)
	} else {
		text.Print("Der Gedanke jemanden in Schmerzen sterben zu lassen gefällt dir mehr als ud erwartet hast...")
		player1.UpdateKarma(-5)
	}
	return player1, inventory, hp, att, def, rec, player_level
}
