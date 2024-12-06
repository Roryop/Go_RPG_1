package gear

import (
	"fmt"
)

type inventorySlot struct {
	itemName string
	count    int
}

func NewInventorySlot() *inventorySlot {
	var slot *inventorySlot = new(inventorySlot)
	return slot
}

func NewInventory() [10]*inventorySlot {
	var inventory [10]*inventorySlot
	fmt.Println(inventory)
	return inventory
}
