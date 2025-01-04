package gear

import "fmt"

type InventorySlot struct {
	item  *Gear
	count int
}

/////////////////// Giving out Stats of Accessoires from Inventory ////////////

func CreateStatsAccessoires(inventory [10]*InventorySlot) (int, int, int) {
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

func GiveInventoryInformation(inventory [10]*InventorySlot) {
	fmt.Println("Inventory:")
	for i := 0; i < 10; i++ {
		fmt.Print("\n Slot ", i+1, ":		", inventory[i].count, " ", inventory[i].item.name)
	}
	fmt.Println()
}

/////////////////// Input to Inventory Slot ///////////////

func (i *InventorySlot) InputInventorySlot(gear *Gear, amount int) {
	i.item = gear
	i.count = amount
}

/////////// Functions for Inventory and Slot Creation /////////

func NewInventorySlot() *InventorySlot {
	var slot *InventorySlot = new(InventorySlot)
	return slot
}

func FillEmptyInventory(inventory [10]*InventorySlot) [10]*InventorySlot {
	for i := 0; i < 10; i++ {
		var emptyItemSlot = NewInventorySlot()
		var emptyItem = NewGear("", "")
		emptyItemSlot.InputInventorySlot(emptyItem, 0)

		inventory[i] = emptyItemSlot
	}

	return inventory
}

func NewInventory() [10]*InventorySlot {
	var inventory [10]*InventorySlot

	inventory = FillEmptyInventory(inventory)

	return inventory
}
