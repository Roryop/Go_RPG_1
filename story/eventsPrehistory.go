package story

import (
	"fmt"
	"start/gear"
	"start/text"
)

////////////////Funktionen für Höhle///////////////////////////

// Gets inventory
// Executes Event HoelenmenschenItem, adjusting Karma or inventory
// Returns inventory
func HoelenmenschenWaffe(inventory [10]*gear.InventorySlot) [10]*gear.InventorySlot {
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

		//////////////////////////////////////////////////////
		//////////////////////////////////////////////////////
		/////////////////////// KARMA + //////////////////////
		//////////////////////////////////////////////////////
		//////////////////////////////////////////////////////

	} else { // Player chooses to NOT gift Item
		text.Print("Die Hölenmenschen werden wegen dir jetzt verhungern!")
		text.Print("Vielleicht sagt dir das etwas...")

		//////////////////////////////////////////////////////
		//////////////////////////////////////////////////////
		/////////////////////// KARMA - //////////////////////
		//////////////////////////////////////////////////////
		//////////////////////////////////////////////////////
	}

	return inventory
}

// Gets hp
// Executes Event HoelenmenschenHunger, adjusting Karma or hp
// Returns hp
func HoelenmenschenHunger(hp int) int {
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

		//////////////////////////////////////////////////////
		//////////////////////////////////////////////////////
		/////////////////////// KARMA - //////////////////////
		//////////////////////////////////////////////////////
		//////////////////////////////////////////////////////
	} else {
		text.Print("Du hast die Hölenmenschen nicht bestohlen.")
		text.Print("Du hungerst...")

		hp = hp - (hp / 10)
	}
	return hp
}

// Gets Nothing
// Executes Event HoelenmenschenWerkzeug, adjusting Karma
// Returns Nothing
func HoelenmenschenWerkzeug() {
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

		//////////////////////////////////////////////////////
		//////////////////////////////////////////////////////
		/////////////////////// KARMA + //////////////////////
		//////////////////////////////////////////////////////
		//////////////////////////////////////////////////////

	} else {
		text.Print("Du bist zu Cool für sowas...")
		text.Print("Du überlässt sie Ihrem eigenen Genie.")
	}
}

////////////////Funktionen für Dschungel//////////////////////////

////////////////Funktionen für Berg///////////////////////////
