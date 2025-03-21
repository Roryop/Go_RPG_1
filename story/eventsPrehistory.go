package story

import (
	"fmt"
	"math/rand"
	"start/gear"
	"start/player"
	"start/text"
)

////////////////Funktionen für Höhle///////////////////////////

// Gets inventory
// Executes Event HoelenmenschenItem, adjusting Karma or inventory
// Returns inventory
func HoehlenmenschenWaffe(player1 *player.Player, inventory [10]*gear.InventorySlot) [10]*gear.InventorySlot {
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
func HoehlenmenschenHunger(player1 *player.Player, hp int) int {
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
func HoehlenmenschenWerkzeug(player1 *player.Player) {
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

// Gets player, inventory, current stats
// player chooses whether to eat woman, if yes -> random permanent stat increase
// Returns current stats
func EatWoman(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int) (int, int, int, int) {
	var choice int

	text.Print("Du triffst auf eine sehr dubiose Frau.")
	text.Space(2)
	text.ShortWait()
	text.Print("Dir grummelt der Magen...")

	text.Print("Möchtest du die Frau essen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	text.Space(5)
	text.LongWait()

	if choice == 1 {
		var randStat int = rand.Intn(4) + 1 //chooses random stat

		text.Print("Die Frau schmeckt sehr lecker.")
		text.ShortWait()
		text.Print("Ein zufälliger Stat hat sich erhöht.")

		hp, att, def, rec = player1.UpgradeStat(inventory, hp, att, def, rec, randStat)

		player1.UpdateKarma(-4) //stat increase at cost of karma
	} else {
		text.Print("Du hast die Frau nicht gegessen...")

		player1.UpdateKarma(1)
	}
	return hp, att, def, rec
}

////////////////Funktionen für Berg///////////////////////////

// Gets player, inventory, current stats, world, player_level
// player chooses path -> wait 15 sec or fight enemy
// Returns player, inventory, current stats, player_level
func MountainPath(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int

	text.Print("Du stehst vor einer Weggabelung.")
	text.Space(2)
	text.ShortWait()

	text.Print("Möchtest du den schnelleren, aber gefährlicheren Pfad gehen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	text.Space(5)
	text.LongWait()

	if choice == 1 {
		text.Print("Du wirst von einer prähistorischen Bergsteigerziege angegriffen")

		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 13)

		text.ShortWait()
		text.Space(5)
		text.Print("Du findest das Nest des Vogels und lootest es.")
		inventory = gear.AddDropToInventory(inventory, player_level)

	} else {
		text.Print("Du wanderst für viele Tage umher...")

		for i := 0; i < 5; i++ {
			text.LongWait()
			text.Print(".")
		}
	}

	return player1, inventory, hp, att, def, rec, player_level
}

// Gets player, inventory, current stats, world, player_level
// Fights a Pterodactylus
// Returns player, inventory, current stats, player_level
func Pterodactylus(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem Pterodactylus angegriffen")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 12)

	text.ShortWait()
	text.Space(5)
	text.Print("Du lootest sein Nest")
	inventory = gear.AddDropToInventory(inventory, player_level)

	return player1, inventory, hp, att, def, rec, player_level
}

// Gets player, inventory, current stats, world, player_level
// Fights Bergsteigerziegen
// Returns player, inventory, current stats, player_level
func Bergsteigerziegen(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	for i := 0; i < 3 && hp > 0; i++ { //Fight gegen 3 Bergsteigerziegen
		text.Print("Du wirst von einer Bergsteigerziege angegriffen")

		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 13)
	}

	text.ShortWait()
	text.Space(5)
	text.Print("Die Ziegen haben eine Schatztruhe beschützt.")
	text.Print("Du öffnest die Schatztruhe.")
	inventory = gear.AddDropToInventory(inventory, player_level)

	return player1, inventory, hp, att, def, rec, player_level
}
