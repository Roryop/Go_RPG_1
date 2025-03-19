package story

import (
	"fmt"
	"start/gear"
	"start/player"
	"start/text"
)

////////////////Funktionen für Wald///////////////////////////

// Spieler wird vom Bär angegriffen
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Baerangriff(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem Bär angegriffen")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 3)

	return player1, inventory, hp, att, def, rec, player_level
}

// Speiler heilt am Lagerfeuer
// Gets PlayerHp, PlayerMaxHp
// Returns PlayerHp mit addiertem ein fünftel vom PlayerMaxHp
func Bonfire(player1 *player.Player, inventory [10]*gear.InventorySlot, hp int) int {
	var maxHp, _, _, _ = player1.CreateStats(inventory)

	text.Print("Du findest einen Lagerfeuer und ruhst dich aus")

	hp = hp + (maxHp / 5)
	if hp > maxHp {
		hp = maxHp
	}

	return hp
}

// Spieler wählt aus ob er einen Pilz isst
// Gets PlayerHP, PlayerMaxHp
// Returns PlayerHp nach essen oder nicht essen vom Pilz
func GesunderPilz(player1 *player.Player, inventory [10]*gear.InventorySlot, hp int) int {
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
		if hp > maxHp {
			hp = maxHp
		}

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
func Hexenjagd(player1 *player.Player) {
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
func Burgritter(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Ein Burgritter fordert dich zu einer Duell heraus.")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 3)

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler wird in den Kerker geworfen
func Kerker(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int) (*player.Player, [10]*gear.InventorySlot, int, int) {
	var choice int
	var maxHp, _, _, _ = player1.CreateStats(inventory)

	text.Print("Du würdest aufgrund eines Verdachts in den Kerker geworfen.")

	// Auswahl Kerker ausbrechen oder nicht
	text.Print("Möchtest du versuchen auszubrechen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Du brichst dir 3 Zehe und 2 Finger.")
		hp = hp - (maxHp / 7)

	} else {
		_, att, _, _ = player1.UpgradeStat(inventory, hp, att, def, rec, 2)
		text.Print("Durch intensives Training und Meditation baust du Masse auf.")
	}
	return player1, inventory, hp, att
}

////////////////Funktionen für Dorf///////////////////////////

// Spieler kämpft gegen Dorfbewohner
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kanpf
func BetrunkenerDorfbewohner(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem betrunkenen Dorfbewohner angegriffen!")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 6)

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler wählt aus ob er steine an den Rotzbuben zurückwirft
// Gets PlayerStats
// returns nothing
func Rotzbuben(player1 *player.Player) {
	var choice int

	text.Print("Du wirst von kleinen Rotzbuben mit Steinen beworfen.")

	// Auswahl Pilz essen oder nicht
	text.Print("Willst du zurückwerfen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Du hast einem Kind den Schädel durchgeschlagen.")
		text.Print("Fühlst du dich jetzt besser?")
	} else {
		text.Print("Du nimmst den Weg des Friedens und die Rotzbuben werden von deren Müttern nach Hause gezogen.")
	}
	///////////////////////////////////////////////
	///////////////////////////////////////////////
	/////////            KARMA           //////////
	///////////////////////////////////////////////
	///////////////////////////////////////////////
}

// Spieler wählt aus ob im Wirtshaus rastet
// Gets PlayerHP, PlayerMaxHp
// Returns PlayerHp nach rasten oder nicht rasten
func Wirtshaus(player1 *player.Player, inventory [10]*gear.InventorySlot, hp int) int {
	var maxHp, _, _, _ = player1.CreateStats(inventory)
	var choice int

	text.Print("Ein Platz im Wirtshaus ist frei.")

	// Auswahl Pilz essen oder nicht
	text.Print("Möchtest du dort essen und rasten?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		hp = hp + (maxHp / 6)
		if hp > maxHp {
			hp = maxHp
		}

		text.Print("Du hast dich geizig voll gefressen.")
	} else {
		text.Print("Du hast auf eine Mahlzeit voller Bio-Produkte verzichtet.")
	}
	return hp
}
