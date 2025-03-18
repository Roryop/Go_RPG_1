package gear

import (
	"fmt"
	"start/text"
)

type InventorySlot struct {
	item  *Gear
	count int
}

/////////////////// Giving out Stats of Items from Inventory ////////////

// Gets Inventory
// Calculates Stats of Items in Inventory
// Returns said Stats (att, def, rec)
func CreateStatsItems(inventory [10]*InventorySlot) (int, int, int) {
	var att int = 0
	var def int = 0
	var rec int = 0

	for i := 0; i < 10; i++ {
		att = att + inventory[i].item.damage*inventory[i].count
		def = def + inventory[i].item.defense*inventory[i].count
		rec = rec + inventory[i].item.recovery*inventory[i].count
	}

	return att, def, rec
}

///////////////////// Adding Item To Inventory /////////////////////

// Gets inventory, world_barrier
// Drops Item and Makes player choose whether to store it (and stores)
// Returns Inventory
func AddDropToInventory(inventory [10]*InventorySlot, player_level int) [10]*InventorySlot {
	var item = ItemDrop(player_level)

	if item.gearTyp != "Empty" {
		text.Print("Der Gegner hat einen " + item.name + " gedroppt!")
		inventory = AddToInventory(inventory, item)
	} else {
		text.Print("Der Gegner hat leider keine Items gedroppt :c")
	}

	return inventory
}

// Gets Inventory + Item
// Puts Item into allocated InventorySlot
// Returns Inventory
func AddToInventory(inventory [10]*InventorySlot, item *Gear) [10]*InventorySlot {
	var wishToStore int

	text.Print("Möchtest du dieses Item in deinem Inventar verstauen?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&wishToStore)

	if wishToStore == 1 {
		var slot int

		for slot < 1 || slot > 10 {
			text.Print("In welchem Slot möchtest du es verstauen?")
			text.ShortWait()

			GiveInventoryInformation(inventory)

			fmt.Scanln(&slot)
		}

		var inventorySlot = NewInventorySlot()
		inventorySlot.InputInventorySlot(item, 1)
		if inventory[slot-1].item.name != inventorySlot.item.name {
			inventory[slot-1] = inventorySlot
		} else {
			inventory[slot-1].count += 1
		}

		text.Print("Das Item wurde verstaut.")
		text.ShortWait()
		text.Print("Überprüfe Inventar:")
		text.ShortWait()

		GiveInventoryInformation(inventory)
	} else {
		text.Print("Du hast dich dazu entschieden, dieses Item nicht zu verstauen.")
		text.Print("Du wirfst es weg.")
	}
	return inventory
}

// Gets Inventory
// Subtracts one Item from Inventory
// Returns Inventory
func SubtractFromInventory(inventory [10]*InventorySlot) [10]*InventorySlot {
	var slot int
	var choice int

	for choice != 1 {
		for slot < 1 || slot > 10 {
			text.Print("Aus welchem Slot möchtest du ein Item aus deinem Inventory entfernen?")
			GiveInventoryInformation(inventory)
			fmt.Scanln(&slot)

			if inventory[slot-1].count < 0 {
				text.Print("In diesem Slot befindet sich kein Item!")
			}
		}

		text.Print("Bist du dir sicher, dass du " + inventory[slot-1].item.name + " entfernen möchtest?")
		fmt.Println("1: Ja")
		fmt.Println("2: Nein")
		fmt.Scanln(&choice)
	}

	if inventory[slot-1].count > 0 {
		inventory[slot-1].count = inventory[slot-1].count - 1 // Count of Item reduced by 1
	}
	if inventory[slot-1].count <= 0 {
		var emptyItemSlot = NewInventorySlot()
		var emptyItem = NewGear("", "")
		emptyItemSlot.InputInventorySlot(emptyItem, 0)

		inventory[slot-1] = emptyItemSlot
	}
	text.Print("Das Item wurde entfernt.")
	text.Print("Checking Inventory...")
	GiveInventoryInformation(inventory)
	return inventory
}

///////////////////// Giving out Inventory Information to user ///////////////

// Gets Inventory
// Gives out Content of all Inventory Slots
// Returns Nothing
func GiveInventoryInformation(inventory [10]*InventorySlot) {
	fmt.Println("Inventar:")
	for i := 0; i < 10; i++ {
		fmt.Print("\n Slot ", i+1, ":		", inventory[i].count, " ", inventory[i].item.name)
	}
	fmt.Println()
}

/////////////////// Input to Inventory Slot ///////////////

// Gets Item of type *Gear + the amount of said Item in current Slot
// Sets Item of current Slot to Item given to Function
// Returns Nothing
func (i *InventorySlot) InputInventorySlot(gear *Gear, amount int) {
	i.item = gear
	i.count = amount
}

/////////// Functions for Inventory and Slot Creation /////////

// Gets Nothing
// Creates InventorySlot
// Returns InventorySlot
func NewInventorySlot() *InventorySlot {
	var slot *InventorySlot = new(InventorySlot)
	return slot
}

// Takes Inventory
// Fills Inventory with Empty Items
// Returns Inventory
func FillEmptyInventory(inventory [10]*InventorySlot) [10]*InventorySlot {
	for i := 0; i < 10; i++ {
		// Creating Empty Item
		var emptyItemSlot = NewInventorySlot()
		var emptyItem = NewGear("", "")
		emptyItemSlot.InputInventorySlot(emptyItem, 0)

		// Filling Current InventorySlot with Empty Item
		inventory[i] = emptyItemSlot
	}

	return inventory
}

// Gets Nothing
// Creates New Inventory + Fills it
// Returns Inventory
func NewInventory() [10]*InventorySlot {
	var inventory [10]*InventorySlot

	inventory = FillEmptyInventory(inventory)

	return inventory
}
