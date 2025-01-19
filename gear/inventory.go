package gear

import "fmt"

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
		att = att + inventory[i].item.damage
		def = def + inventory[i].item.defense
		rec = rec + inventory[i].item.recovery
	}

	return att, def, rec
}

///////////////////// Giving out Inventory Information to user ///////////////

// Gets Inventory
// Gives out Content of all Inventory Slots
// Returns Nothing
func GiveInventoryInformation(inventory [10]*InventorySlot) {
	fmt.Println("Inventory:")
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
