package game

import (
	"fmt"
	"math/rand"
	"start/gear"
)

// Gets Nothing
// Player Chooses Game_level
// Returns Game_level
func ChooseGameLevel() int {
	var game_level = 0
	fmt.Println("In welches Level möchtest du?")
	fmt.Scanln(&game_level)
	fmt.Println("Du hast Level", game_level, "ausgewählt.")
	return game_level
}

// Gets Nothing
// Creates Item based on rand.Intn() + ItemArrays
// Returns Item
func ItemDrop() *gear.Gear {
	var rarity string

	var dropNumber int = rand.Intn(10000) + 1
	switch {
	case dropNumber <= 8000 && dropNumber > 3000:
		rarity = "Common"
	case dropNumber <= 9400 && dropNumber > 8000:
		rarity = "Greater"
	case dropNumber <= 9900 && dropNumber > 9400:
		rarity = "Unique"
	case dropNumber <= 9990 && dropNumber > 9900:
		rarity = "Mythical"
	case dropNumber <= 9999 && dropNumber > 9990:
		rarity = "Transcendent"
	case dropNumber == 10000:
		rarity = "Godly"
	default:
		rarity = "Nothing"
	}

	// Creates Item
	var item *gear.Gear

	// Creates Array for deciding itemTyp
	var itemTypArray [3]string = [3]string{"Weapon", "Armor", "Accessoire"}
	// Deciding on ItemTyp through rand.Intn()
	var itemTypNumber int = rand.Intn(3)
	var itemTyp string = itemTypArray[itemTypNumber]

	// Creating ItemArrays of Typ Common
	var weaponArrayCommon [2]string = [2]string{
		"Dagger(Common)",
		"Spear(Common)"}

	var armorArrayCommon [1]string = [1]string{
		"Leather Breastplate(Common)"}

	var accessoireArrayCommon [3]string = [3]string{
		"Ring of Strength(Common)",
		"Ring of Resilience(Common)",
		"Ring of Recovery(Common)"}

	// Creating Variables to store Item Name of corresponding Item
	var itemName string
	var itemArrayNumber int

	// Assigning ItemName according to rarity + rand.Intn()
	switch rarity {
	case "Nothing":

	case "Common":
		switch itemTyp {
		case "Weapon":
			itemArrayNumber = rand.Intn(len(weaponArrayCommon))
			itemName = weaponArrayCommon[itemArrayNumber]
		case "Armor":
			itemArrayNumber = rand.Intn(len(armorArrayCommon))
			itemName = armorArrayCommon[itemArrayNumber]
		case "Accessoire":
			itemArrayNumber = rand.Intn(len(accessoireArrayCommon))
			itemName = accessoireArrayCommon[itemArrayNumber]
		}

	case "Greater":

	case "Unique":

	case "Mythical":

	case "Transcendent":

	case "Godly":

	}

	// Creating Item based on ItemTyp + ItemName
	item = gear.NewGear(itemTyp, itemName)
	return item
}
