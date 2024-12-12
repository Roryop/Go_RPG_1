package gear

import "fmt"

type inventorySlot struct {
	item  Gear
	count int
}

///////////////////// Giving out Inventory Information to user ///////////////

func GiveInventoryInformation(inventory [10]*inventorySlot) {
	fmt.Println("Inventory:")
	for i := 0; i < 10; i++ {
		fmt.Print("\n Slot ", i+1, ":		", *inventory[0])
	}
	fmt.Println()
}

/////////////////// Input to Inventory Slot ///////////////

func (i *inventorySlot) InputInventorySlot(gear Gear, amount int) {
	i.item = gear
	i.count = amount

}

/////////// Functions for Inventory and Slot Creation /////////

func NewInventorySlot() *inventorySlot {
	var slot *inventorySlot = new(inventorySlot)
	return slot
}

func NewInventory() [10]*inventorySlot {
	var inventory [10]*inventorySlot
	return inventory
}
