package story

import (
	"fmt"
	"start/gear"
	"start/player"
	"start/text"
)

//////////////////////Funktionen für Bar///////////////////////////

// Spieler wählt aus ob er jemanden für Whiskey ausgeben lässt
// Gets PlayerStats
// returns karma
func Whiskey(player1 *player.Player) {
	var choice int

	text.Print("Dir will ein dir unbekannter Einwohner Whiskey ausgeben.")

	// Auswahl Bewohner helfen oder nicht
	text.Print("Nimmst du an?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Der Unbekannte scheint dir sehr charismatisch.")
		text.Print("Whiskey Gläser bringen euch zusammen und ihr freundet euch an.")
		player1.UpdateKarma(7)
	} else {
		text.Print("Aus Höflichkeit oder aus dem Grund, dass der Unbekannte unerträglich stank. Lehnst du ab.")
	}
}

// Spieler wählt aus ob er sich in eine Schlaegerei einmischen will
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Schlaegerei(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int

	text.Print("In der Bar bricht eine Schlaegerei aus!")

	// Auswahl Verteidigung helfen oder nicht
	text.Print("Willst du deich einmischen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Es kommt zu einem Kampf eins gegen eins!")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 14)

	} else {
		text.Print("Du hast dich unter dem Bartisch versteckt. Bis alles vorbei ging.")

	}

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler wählt aus ob er sich in eine Schlaegerei einmischen will
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Barüberfall(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int

	text.Print("Die Bar wird überfallen.")

	// Auswahl Verteidigung helfen oder nicht
	text.Print("Willst du die Bar verteidigen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Der Anführer der Bande taucht vor dir auf...")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 15)
		player1.UpdateKarma(6)
	} else {
		text.Print("Du bist geflüchtet.")
		text.Print("...")
		text.Print("Der Anführer der Bande hatte nicht mal ein Pferd.")
	}

	return player1, inventory, hp, att, def, rec, player_level
}

////////////////Funktionen für Sheriff Büro///////////////////////////

// Spieler wählt aus ob er einen Insassen Freilässt
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Jailbreak(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int
	var item = gear.ItemDrop(player_level)
	var gearTyp = item.GetGearTyp()

	text.Print("Der Insasse des Gefängnisses bittet dich heimlich, ihn freizulassen.")

	// Auswahl Freilassen oder nicht
	text.Print("Willst du auf ihn Hören?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Der Insasse bedankt sich bei dir.")
		if gearTyp == "Empty" {
			text.Print("Jedoch aber nur wörtlich.")
		} else {
			inventory = gear.AddToInventory(inventory, item)
		}
		text.Print("Der Sheriff riecht aber dein Missbrauch von Freiheit von 3/4 Meile aus.")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 16)
	} else {
		text.Print("Du lässt ihn in der Zelle gammeln.")
		player1.UpdateKarma(4)
	}

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler wählt aus ob er auf einen Insassen aufpassen will
// Gets inventory (fügt einen Item hinzu falls der Spieler dem Sheriff hilft)
// returns inventory
func Sheriffquest(player1 *player.Player, inventory [10]*gear.InventorySlot, player_level int) [10]*gear.InventorySlot {
	var choice int
	var item = gear.NewGear("Weapon", "Smith & Wesson 500 Bone Crusher")

	text.Print("Der Sherrif fragt dich, ob du kurz auf einen gefangenen aufpassen kannst.")

	// Auswahl Zeit nehmen oder nicht
	text.Print("Willst du dir die Zeit nehmen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Der Sheriff schenkt dir eine Waffe aus seinem Arsenal.")
		inventory = gear.AddToInventory(inventory, item)
	} else {
		text.Print("Würd mir stinken...")
	}
	return inventory
}

// Spieler wählt aus ob er Sheriffs schubladen durchsuchen will
// Gets inventory (fügt einen Medipack hinzu falls der Spieler den Sheriff beraubt)
// returns inventory
func Sheriffberauben(player1 *player.Player, inventory [10]*gear.InventorySlot, player_level int) [10]*gear.InventorySlot {
	var choice int
	var item = gear.NewGear("Accessoire", "Hello Kitty Verband")

	text.Print("Der Sherrif steckt sich eine Zigarre an und geht raus")

	// Auswahl durchschauen oder nicht
	text.Print("Willst du die Schubladen seines Schreibtisches durchsuchen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Du bist ehrenlos...")
		inventory = gear.AddToInventory(inventory, item)
		player1.UpdateKarma(-4)
	} else {
		text.Print("Guter Junge...")
	}
	return inventory
}

////////////////Funktionen für Goldmine///////////////////////////

// Spieler wählt aus ob er Koyoten angreift
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Koyothoehle(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int
	var item = gear.NewGear("Accessoire", "Actimel")

	text.Print("Du siehst in der Höhle einen Koyoten an der Leiche eines Minenarbeiters nagen.")

	// Auswahl angreifen oder nicht
	text.Print("Möchtest du ihn angreifen und hoffen, dass die Leiche etwas nützliches bei sich trägt?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Du entscheidest dich den Koyoten anzugreifen.")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 17)
		inventory = gear.AddToInventory(inventory, item)
		player1.UpdateKarma(-4)
	} else {
		text.Print("Du entscheidest dich nicht einzumischen.")
	}

	return player1, inventory, hp, att, def, rec, player_level
}

// Spieler kriegt aufs Gesicht wenn er einen Balken bei seite räumt
// Gets PlayerHp, PlayerMaxHp
// Returns PlayerHp mit addiertem ein fünftel vom PlayerMaxHp
func Aufdiefressekriegen(player1 *player.Player, inventory [10]*gear.InventorySlot, hp int) int {
	var maxHp, _, _, _ = player1.CreateStats(inventory)
	var choice int

	text.Print("Hinter einem Balken in der Wand steckt eine stark aussehende Waffe.")

	// Auswahl räumen oder nicht
	text.Print("Willst du ihn beiseite räumen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Ein tragender Balken...")
		text.Print("Ein Teil der Mine stürzt ein, du wirst von einem Stein auf den Kopf getroffen.")
		hp = hp - (maxHp / 7)
		if hp > maxHp {
			hp = maxHp
		}

	} else {
		text.Print("Du erinnerst dich an weisheiten deines Onkels...")
		text.Print("Du bleibst zufrieden mit deiner hetzigen Waffe und gehst gelassen weg...")
	}
	return hp
}

// Spieler wählt aus ob er Banditen überfällt oder nicht
// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach Kampf
func Banditenlager(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int

	text.Print("Du findest einen Lager von Banditen")

	// Auswahl Banditen überfallen oder nicht
	text.Print("Willst du die überfallen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	if choice == 1 {
		text.Print("Der Anführer der Bande taucht vor dir auf...")
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 15)
	} else {
		text.Print("Du hattest kein Bock auf eine Eskalation in einer firedlichen Höhle.")
	}

	return player1, inventory, hp, att, def, rec, player_level
}
