package story

import (
	"fmt"
	"start/gear"
	"start/player"
	"start/text"
)

////////////////Funktionen für Höhle///////////////////////////

// Gets inventory
// Executes Event HoelenmenschenItem, adjusting Karma or inventory
// Returns inventory
func HoelenmenschenWaffe(player1 *player.Player, inventory [10]*gear.InventorySlot) [10]*gear.InventorySlot {
	var choice int

	text.Print("Du triffst auf eine Gruppe Hölenmenschen.")
	text.Space(2)
	text.ShortWait()
	text.Print("Sie haben massive Schwierigkeiten beim Jagen.")

	text.Print("Möchtest du Ihnen eins deiner Items geben?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	text.Space(5)
	text.LongWait()

	if choice == 1 { // Player chooses to gift Item
		inventory = gear.SubtractFromInventory(inventory)

		text.Space(5)
		text.ShortWait()

		text.Print("Die Hölenmenschen können nun dank dir überleben.")
		text.Print("Du scheinst auf dem richtigen Weg zu sein...")

		player1.UpdateKarma(8)

	} else { // Player chooses to NOT gift Item
		text.Print("Die Hölenmenschen werden wegen dir jetzt verhungern!")
		text.Print("Vielleicht sagt dir das etwas...")

		player1.UpdateKarma(-4)
	}

	return inventory
}

// Gets hp
// Executes Event HoelenmenschenHunger, adjusting Karma or hp
// Returns hp
func HoelenmenschenHunger(player1 *player.Player, hp int) int {
	var choice int

	text.Print("Du triffst auf eine Gruppe Hölenmenschen.")
	text.Space(2)
	text.ShortWait()
	text.Print("Sie scheinen massive Vorräte an Nahrung zu haben.")
	text.Print("Du bist dabei zu verhungern...")

	text.Print("Möchtest du ihnen etwas Nahrung stehlen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	text.Space(5)
	text.LongWait()

	if choice == 1 {
		text.Print("Du hast Ihnen einen Teil ihrer Nahrung gestohlen.")
		text.Print("Vielleicht sagt dir das etwas...")

		player1.UpdateKarma(-3)
	} else {
		text.Print("Du hast die Hölenmenschen nicht bestohlen.")
		text.Print("Du hungerst...")

		hp = hp - (hp / 6) // Subtracting 1/6th from current hp
	}
	return hp
}

// Gets Nothing
// Executes Event HoelenmenschenWerkzeug, adjusting Karma
// Returns Nothing
func HoelenmenschenWerkzeug(player1 *player.Player) {
	var choice int

	text.Print("Du triffst auf eine Gruppe Hölenmenschen.")
	text.Space(2)
	text.ShortWait()
	text.Print("Sie scheinen keine Werkzeuge zum Jagen zu besitzen.")

	text.Print("Möchtest du Ihnen beibringen, Werkzeuge zu machen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	text.Space(5)
	text.LongWait()

	if choice == 1 {
		text.Print("Du hast Ihnen beigebracht, Werkzeuge zu machen.")
		text.Print("Du scheinst auf dem richtigen Weg zu sein...")

		player1.UpdateKarma(4)

	} else {
		text.Print("Du bist zu Cool für sowas...")
		text.Print("Du überlässt sie Ihrem eigenen Genie.")
	}
}

////////////////Funktionen für Dschungel//////////////////////////

// Gets player, inventory, current stats, world, player_level
// Fights a Triceratops
// Returns player, inventory, current stats, player_level
func Triceratops(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem Triceratops angegriffen")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 11)

	return player1, inventory, hp, att, def, rec, player_level
}

// Gets player, inventory, current stats, world, player_level
// Takes player to a lake, depending on choice = healing and fight
// Returns player, inventory, current stats, player_level
func Teich(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var maxHp, _, _, _ = player1.CreateStats(inventory)
	var choice int

	text.Print("Du kommst an einem Teich an.")
	text.Space(2)
	text.ShortWait()
	text.Print("Du scheinst ziemlich viel Durst zu haben.")

	text.Print("Möchtest du aus dem Teich trinken?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	text.Space(5)
	text.LongWait()

	if choice == 1 {
		text.Print("Du trinkst aus dem Teich und erholst dich.")

		hp = hp + (maxHp / 5)
		if hp > maxHp {
			hp = maxHp
		}

		text.LongWait()
		text.Print("Du wirst von einem Wolf überrascht")

		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 2)

		player1.UpdateKarma(4)

	} else {
		text.Print("Du trinkst nicht aus dem Teich")
	}

	return player1, inventory, hp, att, def, rec, player_level
}

////////////////Funktionen für Berg///////////////////////////
